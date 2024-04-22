// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/shared/mlog"
	"github.com/mattermost/mattermost/server/v8/channels/audit"
	"github.com/mattermost/mattermost/server/v8/platform/shared/web"
)

func (api *API) InitCloud() {
	// GET /api/v4/cloud/products
	api.BaseRoutes.Cloud.Handle("/products", api.APISessionRequired(getCloudProducts)).Methods("GET")
	// GET /api/v4/cloud/limits
	api.BaseRoutes.Cloud.Handle("/limits", api.APISessionRequired(getCloudLimits)).Methods("GET")

	api.BaseRoutes.Cloud.Handle("/products/selfhosted", api.APISessionRequired(getSelfHostedProducts)).Methods("GET")

	// POST /api/v4/cloud/payment
	// POST /api/v4/cloud/payment/confirm
	api.BaseRoutes.Cloud.Handle("/payment", api.APISessionRequired(createCustomerPayment)).Methods("POST")
	api.BaseRoutes.Cloud.Handle("/payment/confirm", api.APISessionRequired(confirmCustomerPayment)).Methods("POST")

	// GET /api/v4/cloud/customer
	// PUT /api/v4/cloud/customer
	// PUT /api/v4/cloud/customer/address
	api.BaseRoutes.Cloud.Handle("/customer", api.APISessionRequired(getCloudCustomer)).Methods("GET")
	api.BaseRoutes.Cloud.Handle("/customer", api.APISessionRequired(updateCloudCustomer)).Methods("PUT")
	api.BaseRoutes.Cloud.Handle("/customer/address", api.APISessionRequired(updateCloudCustomerAddress)).Methods("PUT")

	// GET /api/v4/cloud/subscription
	api.BaseRoutes.Cloud.Handle("/subscription", api.APISessionRequired(getSubscription)).Methods("GET")
	api.BaseRoutes.Cloud.Handle("/subscription/invoices", api.APISessionRequired(getInvoicesForSubscription)).Methods("GET")
	api.BaseRoutes.Cloud.Handle("/subscription/invoices/{invoice_id:[_A-Za-z0-9]+}/pdf", api.APISessionRequired(getSubscriptionInvoicePDF)).Methods("GET")
	api.BaseRoutes.Cloud.Handle("/subscription/self-serve-status", api.APISessionRequired(getLicenseSelfServeStatus)).Methods("GET")
	api.BaseRoutes.Cloud.Handle("/subscription", api.APISessionRequired(changeSubscription)).Methods("PUT")

	// GET /api/v4/cloud/request-trial
	api.BaseRoutes.Cloud.Handle("/request-trial", api.APISessionRequired(requestCloudTrial)).Methods("PUT")

	// GET /api/v4/cloud/validate-business-email
	api.BaseRoutes.Cloud.Handle("/validate-business-email", api.APISessionRequired(validateBusinessEmail)).Methods("POST")
	api.BaseRoutes.Cloud.Handle("/validate-workspace-business-email", api.APISessionRequired(validateWorkspaceBusinessEmail)).Methods("POST")

	// POST /api/v4/cloud/webhook
	api.BaseRoutes.Cloud.Handle("/webhook", api.CloudAPIKeyRequired(handleCWSWebhook)).Methods("POST")

	// GET /api/v4/cloud/installation
	api.BaseRoutes.Cloud.Handle("/installation", api.APISessionRequired(getInstallation)).Methods("GET")

	// GET /api/v4/cloud/cws-health-check
	api.BaseRoutes.Cloud.Handle("/check-cws-connection", api.APIHandler(handleCheckCWSConnection)).Methods("GET")

	api.BaseRoutes.Cloud.Handle("/delete-workspace", api.APISessionRequired(selfServeDeleteWorkspace)).Methods("DELETE")
}

func ensureCloudInterface(c *Context, where string) bool {
	cloud := c.App.Cloud()
	disabled := c.App.Config().CloudSettings.Disable != nil && *c.App.Config().CloudSettings.Disable
	if cloud == nil {
		c.Err = model.NewAppError(where, "api.server.cws.needs_enterprise_edition", nil, "", http.StatusBadRequest)
		return false
	}
	if disabled {
		c.Err = model.NewAppError(where, "api.server.cws.disabled", nil, "", http.StatusUnprocessableEntity)
		return false
	}
	return true
}

func getSubscription(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getSubscription")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.getSubscription", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	subscription, err := c.App.Cloud().GetSubscription(c.AppContext.Session().UserId)
	if err != nil {
		c.Err = model.NewAppError("Api4.getSubscription", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	// if it is an end user, return basic subscription data without sensitive information
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleReadBilling) {
		subscription = &model.Subscription{
			ID:              subscription.ID,
			ProductID:       subscription.ProductID,
			IsFreeTrial:     subscription.IsFreeTrial,
			TrialEndAt:      subscription.TrialEndAt,
			EndAt:           subscription.EndAt,
			CancelAt:        subscription.CancelAt,
			DelinquentSince: subscription.DelinquentSince,
			CustomerID:      "",
			AddOns:          []string{},
			StartAt:         0,
			CreateAt:        0,
			Seats:           0,
			Status:          "",
			DNS:             "",
			LastInvoice:     &model.Invoice{},
			BillingType:     "",
		}
	}

	if model.GetServiceEnvironment() != model.ServiceEnvironmentTest {
		subscription.SimulatedCurrentTimeMs = nil
	}

	if !c.App.Config().FeatureFlags.CloudAnnualRenewals {
		subscription.WillRenew = ""
		subscription.CancelAt = nil
	}

	json, err := json.Marshal(subscription)
	if err != nil {
		c.Err = model.NewAppError("Api4.getSubscription", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	w.Write(json)
}

func changeSubscription(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.changeSubscription")
	if !ensured {
		return
	}
	userId := c.AppContext.Session().UserId

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.changeSubscription", "api.cloud.license_error", nil, "", http.StatusInternalServerError)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleWriteBilling) {
		c.SetPermissionError(model.PermissionSysconsoleWriteBilling)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.changeSubscription", "api.cloud.app_error", nil, "", http.StatusBadRequest).Wrap(err)
		return
	}

	var subscriptionChange *model.SubscriptionChange
	if err = json.Unmarshal(bodyBytes, &subscriptionChange); err != nil || subscriptionChange == nil {
		c.Err = model.NewAppError("Api4.changeSubscription", "api.cloud.app_error", nil, "", http.StatusBadRequest).Wrap(err)
		return
	}

	currentSubscription, appErr := c.App.Cloud().GetSubscription(userId)
	if appErr != nil {
		c.Err = model.NewAppError("Api4.changeSubscription", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
		return
	}

	changedSub, err := c.App.Cloud().ChangeSubscription(userId, currentSubscription.ID, subscriptionChange)
	if err != nil {
		appErr := model.NewAppError("Api4.changeSubscription", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		if err.Error() == "compliance-failed" {
			c.Logger.Error("Compliance check failed", mlog.Err(err))
			appErr.StatusCode = http.StatusUnprocessableEntity
		}

		c.Err = appErr
		return
	}

	if subscriptionChange.Feedback != nil {
		c.App.Srv().GetTelemetryService().SendTelemetry("downgrade_feedback", subscriptionChange.Feedback.ToMap())
	}

	json, err := json.Marshal(changedSub)
	if err != nil {
		c.Err = model.NewAppError("Api4.changeSubscription", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	product, err := c.App.Cloud().GetCloudProduct(c.AppContext.Session().UserId, subscriptionChange.ProductID)
	if err != nil || product == nil {
		c.Logger.Error("Error finding the new cloud product", mlog.Err(err))
	}

	if product.SKU == string(model.SkuCloudStarter) {
		w.Write(json)
		return
	}

	isYearly := product.IsYearly()

	// Log failures for purchase confirmation email, but don't show an error to the user so as not to confuse them
	// At this point, the upgrade is complete.
	if appErr := c.App.SendUpgradeConfirmationEmail(isYearly); appErr != nil {
		c.Logger.Error("Error sending purchase confirmation email", mlog.Err(appErr))
	}

	w.Write(json)
}

func requestCloudTrial(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.requestCloudTrial")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.requestCloudTrial", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleWriteBilling) {
		c.SetPermissionError(model.PermissionSysconsoleWriteBilling)
		return
	}

	// check if the email needs to be set
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.requestCloudTrial", "api.cloud.app_error", nil, "", http.StatusBadRequest).Wrap(err)
		return
	}
	// this value will not be empty when both emails (user admin and CWS customer) are not business email and
	// a new business email was provided via the request business email modal
	var startTrialRequest *model.StartCloudTrialRequest
	if err = json.Unmarshal(bodyBytes, &startTrialRequest); err != nil || startTrialRequest == nil {
		c.Err = model.NewAppError("Api4.requestCloudTrial", "api.cloud.app_error", nil, "", http.StatusBadRequest).Wrap(err)
		return
	}

	changedSub, err := c.App.Cloud().RequestCloudTrial(c.AppContext.Session().UserId, startTrialRequest.SubscriptionID, startTrialRequest.Email)
	if err != nil {
		c.Err = model.NewAppError("Api4.requestCloudTrial", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	json, err := json.Marshal(changedSub)
	if err != nil {
		c.Err = model.NewAppError("Api4.requestCloudTrial", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	defer c.App.Srv().Cloud.InvalidateCaches()

	w.Write(json)
}

func validateBusinessEmail(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.validateBusinessEmail")
	if !ensured {
		return
	}

	user, appErr := c.App.GetUser(c.AppContext.Session().UserId)
	if appErr != nil {
		c.Err = model.NewAppError("Api4.validateBusinessEmail", "api.cloud.request_error", nil, "", http.StatusForbidden).Wrap(appErr)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.requestCloudTrial", "api.cloud.app_error", nil, "", http.StatusBadRequest).Wrap(err)
		return
	}

	var emailToValidate *model.ValidateBusinessEmailRequest
	err = json.Unmarshal(bodyBytes, &emailToValidate)
	if err != nil || emailToValidate == nil {
		c.Err = model.NewAppError("Api4.requestCloudTrial", "api.cloud.app_error", nil, "", http.StatusBadRequest).Wrap(err)
		return
	}

	err = c.App.Cloud().ValidateBusinessEmail(user.Id, emailToValidate.Email)
	if err != nil {
		c.Err = model.NewAppError("Api4.validateBusinessEmail", "api.cloud.request_error", nil, "", http.StatusForbidden).Wrap(err)
		emailResp := model.ValidateBusinessEmailResponse{IsValid: false}
		if err := json.NewEncoder(w).Encode(emailResp); err != nil {
			c.Logger.Warn("Error while writing response", mlog.Err(err))
		}
		return
	}
	emailResp := model.ValidateBusinessEmailResponse{IsValid: true}
	if err := json.NewEncoder(w).Encode(emailResp); err != nil {
		c.Logger.Warn("Error while writing response", mlog.Err(err))
	}
}

func validateWorkspaceBusinessEmail(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.validateWorkspaceBusinessEmail")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.validateWorkspaceBusinessEmail", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleWriteBilling) {
		c.SetPermissionError(model.PermissionSysconsoleWriteBilling)
		return
	}

	user, userErr := c.App.GetUser(c.AppContext.Session().UserId)
	if userErr != nil {
		c.Err = userErr
		return
	}

	// get the cloud customer email to validate if is a valid business email
	cloudCustomer, err := c.App.Cloud().GetCloudCustomer(user.Id)
	if err != nil {
		c.Err = model.NewAppError("Api4.validateWorkspaceBusinessEmail", "api.cloud.request_error", nil, "", http.StatusBadRequest).Wrap(err)
		return
	}
	emailErr := c.App.Cloud().ValidateBusinessEmail(user.Id, cloudCustomer.Email)

	// if the current workspace email is not a valid business email
	if emailErr != nil {
		// grab the current admin email and validate it
		errValidatingAdminEmail := c.App.Cloud().ValidateBusinessEmail(user.Id, user.Email)
		if errValidatingAdminEmail != nil {
			c.Err = model.NewAppError("Api4.validateWorkspaceBusinessEmail", "api.cloud.request_error", nil, "", http.StatusForbidden).Wrap(errValidatingAdminEmail)
			emailResp := model.ValidateBusinessEmailResponse{IsValid: false}
			if err := json.NewEncoder(w).Encode(emailResp); err != nil {
				c.Logger.Warn("Error while writing response", mlog.Err(err))
			}
			return
		}
	}

	// if any of the emails is valid, return ok
	emailResp := model.ValidateBusinessEmailResponse{IsValid: true}
	if err := json.NewEncoder(w).Encode(emailResp); err != nil {
		c.Logger.Warn("Error while writing response", mlog.Err(err))
	}
}

func getSelfHostedProducts(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getSelfHostedProducts")
	if !ensured {
		return
	}

	products, err := c.App.Cloud().GetSelfHostedProducts(c.AppContext.Session().UserId)
	if err != nil {
		c.Err = model.NewAppError("Api4.getSelfHostedProducts", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	byteProductsData, err := json.Marshal(products)
	if err != nil {
		c.Err = model.NewAppError("Api4.getSelfHostedProducts", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleReadBilling) {
		sanitizedProducts := []model.UserFacingProduct{}
		err = json.Unmarshal(byteProductsData, &sanitizedProducts)
		if err != nil || sanitizedProducts == nil {
			c.Err = model.NewAppError("Api4.getSelfHostedProducts", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
			return
		}

		byteSanitizedProductsData, err := json.Marshal(sanitizedProducts)
		if err != nil {
			c.Err = model.NewAppError("Api4.getSelfHostedProducts", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
			return
		}

		w.Write(byteSanitizedProductsData)
		return
	}

	w.Write(byteProductsData)
}

func getCloudProducts(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getCloudProducts")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.getCloudProducts", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	includeLegacyProducts := r.URL.Query().Get("include_legacy") == "true"

	products, err := c.App.Cloud().GetCloudProducts(c.AppContext.Session().UserId, includeLegacyProducts)
	if err != nil {
		c.Err = model.NewAppError("Api4.getCloudProducts", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	byteProductsData, err := json.Marshal(products)
	if err != nil {
		c.Err = model.NewAppError("Api4.getCloudProducts", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleReadBilling) {
		sanitizedProducts := []model.UserFacingProduct{}
		err = json.Unmarshal(byteProductsData, &sanitizedProducts)
		if err != nil {
			c.Err = model.NewAppError("Api4.getCloudProducts", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
			return
		}

		byteSanitizedProductsData, err := json.Marshal(sanitizedProducts)
		if err != nil {
			c.Err = model.NewAppError("Api4.getCloudProducts", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
			return
		}

		w.Write(byteSanitizedProductsData)
		return
	}

	w.Write(byteProductsData)
}

func getCloudLimits(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getCloudLimits")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.getCloudLimits", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	limits, err := c.App.Cloud().GetCloudLimits(c.AppContext.Session().UserId)
	if err != nil {
		c.Err = model.NewAppError("Api4.getCloudLimits", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	json, err := json.Marshal(limits)
	if err != nil {
		c.Err = model.NewAppError("Api4.getCloudLimits", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	w.Write(json)
}

func getCloudCustomer(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getCloudCustomer")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.getCloudCustomer", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleReadBilling) {
		c.SetPermissionError(model.PermissionSysconsoleReadBilling)
		return
	}

	customer, err := c.App.Cloud().GetCloudCustomer(c.AppContext.Session().UserId)
	if err != nil {
		c.Err = model.NewAppError("Api4.getCloudCustomer", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	json, err := json.Marshal(customer)
	if err != nil {
		c.Err = model.NewAppError("Api4.getCloudCustomer", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	w.Write(json)
}

func getInstallation(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getInstallation")
	if !ensured {
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleReadIPFilters) {
		c.SetPermissionError(model.PermissionSysconsoleReadIPFilters)
		return
	}

	installation, err := c.App.Cloud().GetInstallation(c.AppContext.Session().UserId)
	if err != nil {
		c.Err = model.NewAppError("Api4.getInstallation", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	if err := json.NewEncoder(w).Encode(installation); err != nil {
		c.Err = model.NewAppError("Api4.getInstallation", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}
}

// getLicenseSelfServeStatus makes check for the license in the CWS self-serve portal and establishes if the license is renewable, expandable etc.
func getLicenseSelfServeStatus(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getLicenseSelfServeStatus")
	if !ensured {
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageLicenseInformation) {
		c.SetPermissionError(model.PermissionManageLicenseInformation)
		return
	}

	_, token, err := c.App.Srv().GenerateLicenseRenewalLink()

	if err != nil {
		c.Err = err
		return
	}

	status, cloudErr := c.App.Cloud().GetLicenseSelfServeStatus(c.AppContext.Session().UserId, token)
	if cloudErr != nil {
		c.Err = model.NewAppError("Api4.getLicenseSelfServeStatus", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(cloudErr)
		return
	}

	json, jsonErr := json.Marshal(status)
	if jsonErr != nil {
		c.Err = model.NewAppError("Api4.getLicenseSelfServeStatus", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(jsonErr)
		return
	}

	w.Write(json)
}

func updateCloudCustomer(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.updateCloudCustomer")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.updateCloudCustomer", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleWriteBilling) {
		c.SetPermissionError(model.PermissionSysconsoleWriteBilling)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.updateCloudCustomer", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	var customerInfo *model.CloudCustomerInfo
	if err = json.Unmarshal(bodyBytes, &customerInfo); err != nil || customerInfo == nil {
		c.Err = model.NewAppError("Api4.updateCloudCustomer", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	customer, appErr := c.App.Cloud().UpdateCloudCustomer(c.AppContext.Session().UserId, customerInfo)
	if appErr != nil {
		c.Err = model.NewAppError("Api4.updateCloudCustomer", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
		return
	}

	json, err := json.Marshal(customer)
	if err != nil {
		c.Err = model.NewAppError("Api4.updateCloudCustomer", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	w.Write(json)
}

func updateCloudCustomerAddress(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.updateCloudCustomerAddress")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.updateCloudCustomerAddress", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleWriteBilling) {
		c.SetPermissionError(model.PermissionSysconsoleWriteBilling)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.updateCloudCustomerAddress", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	var address *model.Address
	if err = json.Unmarshal(bodyBytes, &address); err != nil || address == nil {
		c.Err = model.NewAppError("Api4.updateCloudCustomerAddress", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	customer, appErr := c.App.Cloud().UpdateCloudCustomerAddress(c.AppContext.Session().UserId, address)
	if appErr != nil {
		c.Err = model.NewAppError("Api4.updateCloudCustomerAddress", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
		return
	}

	json, err := json.Marshal(customer)
	if err != nil {
		c.Err = model.NewAppError("Api4.updateCloudCustomerAddress", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	w.Write(json)
}

func createCustomerPayment(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.createCustomerPayment")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.createCustomerPayment", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleWriteBilling) {
		c.SetPermissionError(model.PermissionSysconsoleWriteBilling)
		return
	}

	auditRec := c.MakeAuditRecord("createCustomerPayment", audit.Fail)
	defer c.LogAuditRec(auditRec)

	intent, err := c.App.Cloud().CreateCustomerPayment(c.AppContext.Session().UserId)
	if err != nil {
		c.Err = model.NewAppError("Api4.createCustomerPayment", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	json, err := json.Marshal(intent)
	if err != nil {
		c.Err = model.NewAppError("Api4.createCustomerPayment", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	auditRec.Success()

	w.Write(json)
}

func confirmCustomerPayment(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.confirmCustomerPayment")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.confirmCustomerPayment", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleWriteBilling) {
		c.SetPermissionError(model.PermissionSysconsoleWriteBilling)
		return
	}

	auditRec := c.MakeAuditRecord("confirmCustomerPayment", audit.Fail)
	defer c.LogAuditRec(auditRec)

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.confirmCustomerPayment", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	var confirmRequest *model.ConfirmPaymentMethodRequest
	if err = json.Unmarshal(bodyBytes, &confirmRequest); err != nil || confirmRequest == nil {
		c.Err = model.NewAppError("Api4.confirmCustomerPayment", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	err = c.App.Cloud().ConfirmCustomerPayment(c.AppContext.Session().UserId, confirmRequest)
	if err != nil {
		c.Err = model.NewAppError("Api4.createCustomerPayment", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	auditRec.Success()

	ReturnStatusOK(w)
}

func getInvoicesForSubscription(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getInvoicesForSubscription")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.getInvoicesForSubscription", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleReadBilling) {
		c.SetPermissionError(model.PermissionSysconsoleReadBilling)
		return
	}

	invoices, appErr := c.App.Cloud().GetInvoicesForSubscription(c.AppContext.Session().UserId)
	if appErr != nil {
		c.Err = model.NewAppError("Api4.getInvoicesForSubscription", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
		return
	}

	json, err := json.Marshal(invoices)
	if err != nil {
		c.Err = model.NewAppError("Api4.getInvoicesForSubscription", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	w.Write(json)
}

func getSubscriptionInvoicePDF(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.getSubscriptionInvoicePDF")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.getSubscriptionInvoicePDF", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	c.RequireInvoiceId()
	if c.Err != nil {
		return
	}

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleReadBilling) {
		c.SetPermissionError(model.PermissionSysconsoleReadBilling)
		return
	}

	pdfData, filename, appErr := c.App.Cloud().GetInvoicePDF(c.AppContext.Session().UserId, c.Params.InvoiceId)
	if appErr != nil {
		c.Err = model.NewAppError("Api4.getSubscriptionInvoicePDF", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(appErr)
		return
	}

	web.WriteFileResponse(
		filename,
		"application/pdf",
		int64(binary.Size(pdfData)),
		time.Now(),
		*c.App.Config().ServiceSettings.WebserverMode,
		bytes.NewReader(pdfData),
		false,
		w,
		r,
	)
}

func handleCWSWebhook(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.handleCWSWebhook")
	if !ensured {
		return
	}

	if !c.App.Channels().License().IsCloud() {
		c.Err = model.NewAppError("Api4.handleCWSWebhook", "api.cloud.license_error", nil, "", http.StatusForbidden)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.handleCWSWebhook", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}
	defer r.Body.Close()

	var event *model.CWSWebhookPayload
	if err = json.Unmarshal(bodyBytes, &event); err != nil || event == nil {
		c.Err = model.NewAppError("Api4.handleCWSWebhook", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	switch event.Event {
	case model.EventTypeFailedPayment:
		if nErr := c.App.SendPaymentFailedEmail(event.FailedPayment); nErr != nil {
			c.Err = nErr
			return
		}
	case model.EventTypeFailedPaymentNoCard:
		if nErr := c.App.SendNoCardPaymentFailedEmail(); nErr != nil {
			c.Err = nErr
			return
		}
	case model.EventTypeSendUpgradeConfirmationEmail:

		// isYearly determines whether to send the yearly or monthly Upgrade email
		isYearly := false
		if event.Subscription != nil && event.CloudWorkspaceOwner != nil {
			user, appErr := c.App.GetUserByUsername(event.CloudWorkspaceOwner.UserName)
			if appErr != nil {
				c.Err = model.NewAppError("Api4.handleCWSWebhook", appErr.Id, nil, "", appErr.StatusCode).Wrap(appErr)
				return
			}

			// Get the current cloud product to determine whether it's a monthly or yearly product
			product, err := c.App.Cloud().GetCloudProduct(user.Id, event.Subscription.ProductID)
			if err != nil {
				c.Err = model.NewAppError("Api4.handleCWSWebhook", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
				return
			}
			isYearly = product.IsYearly()
		}

		if nErr := c.App.SendUpgradeConfirmationEmail(isYearly); nErr != nil {
			c.Err = nErr
			return
		}
	case model.EventTypeSendAdminWelcomeEmail:
		user, appErr := c.App.GetUserByUsername(event.CloudWorkspaceOwner.UserName)
		if appErr != nil {
			c.Err = model.NewAppError("Api4.handleCWSWebhook", appErr.Id, nil, "", appErr.StatusCode).Wrap(appErr)
			return
		}

		teams, appErr := c.App.GetAllTeams()
		if appErr != nil {
			c.Err = model.NewAppError("Api4.handleCWSWebhook", appErr.Id, nil, "", appErr.StatusCode).Wrap(appErr)
			return
		}

		team := teams[0]

		subscription, err := c.App.Cloud().GetSubscription(user.Id)
		if err != nil {
			c.Err = model.NewAppError("Api4.handleCWSWebhook", "api.cloud.request_error", nil, "", http.StatusInternalServerError).Wrap(err)
			return
		}

		if err := c.App.Srv().EmailService.SendCloudWelcomeEmail(user.Email, user.Locale, team.InviteId, subscription.GetWorkSpaceNameFromDNS(), subscription.DNS, *c.App.Config().ServiceSettings.SiteURL); err != nil {
			c.Err = model.NewAppError("SendCloudWelcomeEmail", "api.user.send_cloud_welcome_email.error", nil, "", http.StatusInternalServerError).Wrap(err)
			return
		}
	case model.EventTypeTriggerDelinquencyEmail:
		var emailToTrigger model.DelinquencyEmail
		if event.DelinquencyEmail != nil {
			emailToTrigger = model.DelinquencyEmail(event.DelinquencyEmail.EmailToTrigger)
		} else {
			c.Err = model.NewAppError("Api4.handleCWSWebhook", "api.cloud.delinquency_email.missing_email_to_trigger", nil, "", http.StatusInternalServerError)
			return
		}
		if nErr := c.App.SendDelinquencyEmail(emailToTrigger); nErr != nil {
			c.Err = nErr
			return
		}

	default:
		c.Err = model.NewAppError("Api4.handleCWSWebhook", "api.cloud.cws_webhook_event_missing_error", nil, "", http.StatusNotFound)
		return
	}

	ReturnStatusOK(w)
}

func handleCheckCWSConnection(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.handleCheckCWSConnection")
	if !ensured {
		return
	}

	if err := c.App.Cloud().CheckCWSConnection(c.AppContext.Session().UserId); err != nil {
		c.Err = model.NewAppError("Api4.handleCWSHealthCheck", "api.server.cws.health_check.app_error", nil, "CWS Server is not available.", http.StatusInternalServerError)
		return
	}

	ReturnStatusOK(w)
}

func selfServeDeleteWorkspace(c *Context, w http.ResponseWriter, r *http.Request) {
	ensured := ensureCloudInterface(c, "Api4.selfServeDeleteWorkspace")
	if !ensured {
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		c.Err = model.NewAppError("Api4.selfServeDeleteWorkspace", "api.cloud.app_error", nil, "", http.StatusBadRequest).Wrap(err)
		return
	}
	defer r.Body.Close()

	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionSysconsoleWriteBilling) {
		c.SetPermissionError(model.PermissionSysconsoleWriteBilling)
		return
	}

	var deleteRequest *model.WorkspaceDeletionRequest
	if err = json.Unmarshal(bodyBytes, &deleteRequest); err != nil || deleteRequest == nil {
		c.Err = model.NewAppError("Api4.selfServeDeleteWorkspace", "api.cloud.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		return
	}

	if err := c.App.Cloud().SelfServeDeleteWorkspace(c.AppContext.Session().UserId, deleteRequest); err != nil {
		c.Err = model.NewAppError("Api4.selfServeDeleteWorkspace", "api.server.cws.delete_workspace.app_error", nil, "CWS Server failed to delete workspace.", http.StatusInternalServerError)
		return
	}

	c.App.Srv().GetTelemetryService().SendTelemetry("delete_workspace_feedback", deleteRequest.Feedback.ToMap())

	ReturnStatusOK(w)
}

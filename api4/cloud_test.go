// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mattermost/mattermost-server/v6/einterfaces/mocks"
	"github.com/mattermost/mattermost-server/v6/model"
)

func Test_getCloudLimits(t *testing.T) {
	t.Run("no license returns not implemented", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.App.Srv().RemoveLicense()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		limits, r, err := th.Client.GetProductLimits()
		require.Error(t, err)
		require.Nil(t, limits)
		require.Equal(t, http.StatusForbidden, r.StatusCode, "Expected 403 forbidden")
	})

	t.Run("non cloud license returns not implemented", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.App.Srv().SetLicense(model.NewTestLicense())

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		limits, r, err := th.Client.GetProductLimits()
		require.Error(t, err)
		require.Nil(t, limits)
		require.Equal(t, http.StatusForbidden, r.StatusCode, "Expected 403 forbidden")
	})

	t.Run("error fetching limits returns internal server error", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := &mocks.CloudInterface{}
		cloud.Mock.On("GetCloudLimits", mock.Anything).Return(nil, errors.New("Unable to get limits"))

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = cloud

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		limits, r, err := th.Client.GetProductLimits()
		require.Error(t, err)
		require.Nil(t, limits)
		require.Equal(t, http.StatusInternalServerError, r.StatusCode, "Expected 500 Internal Server Error")
	})

	t.Run("unauthenticated users can not access", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Logout()

		limits, r, err := th.Client.GetProductLimits()
		require.Error(t, err)
		require.Nil(t, limits)
		require.Equal(t, http.StatusUnauthorized, r.StatusCode, "Expected 401 Unauthorized")
	})

	t.Run("good request with cloud server", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := &mocks.CloudInterface{}
		ten := 10
		mockLimits := &model.ProductLimits{
			Messages: &model.MessagesLimits{
				History: &ten,
			},
		}
		cloud.Mock.On("GetCloudLimits", mock.Anything).Return(mockLimits, nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = cloud

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		limits, r, err := th.Client.GetProductLimits()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, r.StatusCode, "Expected 200 OK")
		require.Equal(t, mockLimits, limits)
		require.Equal(t, *mockLimits.Messages.History, *limits.Messages.History)
	})
}

func Test_GetSubscription(t *testing.T) {
	subscription := &model.Subscription{
		ID:          "MySubscriptionID",
		CustomerID:  "MyCustomer",
		ProductID:   "SomeProductId",
		AddOns:      []string{},
		StartAt:     1000000000,
		EndAt:       2000000000,
		CreateAt:    1000000000,
		Seats:       10,
		IsFreeTrial: "true",
		DNS:         "some.dns.server",
		IsPaidTier:  "false",
		TrialEndAt:  2000000000,
		LastInvoice: &model.Invoice{},
	}

	userFacingSubscription := &model.Subscription{
		ID:          "MySubscriptionID",
		CustomerID:  "",
		ProductID:   "SomeProductId",
		AddOns:      []string{},
		StartAt:     0,
		EndAt:       0,
		CreateAt:    0,
		Seats:       0,
		IsFreeTrial: "true",
		DNS:         "",
		IsPaidTier:  "",
		TrialEndAt:  2000000000,
		LastInvoice: &model.Invoice{},
	}

	t.Run("NON Admin users receive the user facing subscription", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("GetSubscription", mock.Anything).Return(subscription, nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		subscriptionReturned, r, err := th.Client.GetSubscription()

		require.NoError(t, err)
		require.Equal(t, subscriptionReturned, userFacingSubscription)
		require.Equal(t, http.StatusOK, r.StatusCode, "Status OK")
	})

	t.Run("Admin users receive the full subscription information", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("GetSubscription", mock.Anything).Return(subscription, nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		subscriptionReturned, r, err := th.SystemAdminClient.GetSubscription()

		require.NoError(t, err)
		require.Equal(t, subscriptionReturned, subscription)
		require.Equal(t, http.StatusOK, r.StatusCode, "Status OK")
	})
}

func Test_requestTrial(t *testing.T) {
	subscription := &model.Subscription{
		ID:         "MySubscriptionID",
		CustomerID: "MyCustomer",
		ProductID:  "SomeProductId",
		AddOns:     []string{},
		StartAt:    1000000000,
		EndAt:      2000000000,
		CreateAt:   1000000000,
		Seats:      10,
		DNS:        "some.dns.server",
		IsPaidTier: "false",
	}

	newValidBusinessEmail := model.StartCloudTrialRequest{Email: ""}

	t.Run("NON Admin users are UNABLE to request the trial", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("GetSubscription", mock.Anything).Return(subscription, nil)
		cloud.Mock.On("RequestCloudTrial", mock.Anything, mock.Anything, "").Return(subscription, nil)
		cloud.Mock.On("InvalidateCaches").Return(nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		subscriptionChanged, r, err := th.Client.RequestCloudTrial(&newValidBusinessEmail)
		require.Error(t, err)
		require.Nil(t, subscriptionChanged)
		require.Equal(t, http.StatusForbidden, r.StatusCode, "403 Forbidden")
	})

	t.Run("ADMIN user are ABLE to request the trial", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("GetSubscription", mock.Anything).Return(subscription, nil)
		cloud.Mock.On("RequestCloudTrial", mock.Anything, mock.Anything, "").Return(subscription, nil)
		cloud.Mock.On("InvalidateCaches").Return(nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		subscriptionChanged, r, err := th.SystemAdminClient.RequestCloudTrial(&newValidBusinessEmail)

		require.NoError(t, err)
		require.Equal(t, subscriptionChanged, subscription)
		require.Equal(t, http.StatusOK, r.StatusCode, "Status OK")
	})

	t.Run("ADMIN user are ABLE to request the trial with valid business email", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		// patch the customer with the additional contact updated with the valid business email
		newValidBusinessEmail.Email = *model.NewString("valid.email@mattermost.com")

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("GetSubscription", mock.Anything).Return(subscription, nil)
		cloud.Mock.On("RequestCloudTrial", mock.Anything, mock.Anything, "valid.email@mattermost.com").Return(subscription, nil)
		cloud.Mock.On("InvalidateCaches").Return(nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		subscriptionChanged, r, err := th.SystemAdminClient.RequestCloudTrial(&newValidBusinessEmail)

		require.NoError(t, err)
		require.Equal(t, subscriptionChanged, subscription)
		require.Equal(t, http.StatusOK, r.StatusCode, "Status OK")
	})
}

func TestNotifyAdmin(t *testing.T) {
	t.Run("error when plan is unknown when notifying on upgrade", func(t *testing.T) {
		th := Setup(t).InitBasic().InitLogin()
		defer th.TearDown()

		statusCode, err := th.Client.NotifyAdmin(&model.NotifyAdminToUpgradeRequest{
			RequiredPlan:    "Unknown plan",
			RequiredFeature: model.PaidFeatureAllProfessionalfeatures,
		})

		require.Error(t, err)
		require.Equal(t, err.Error(), ": Unable to save notify data.")
		require.Equal(t, http.StatusInternalServerError, statusCode)

	})

	t.Run("error when plan is unknown when notifying to trial", func(t *testing.T) {
		th := Setup(t).InitBasic().InitLogin()
		defer th.TearDown()

		statusCode, err := th.Client.NotifyAdmin(&model.NotifyAdminToUpgradeRequest{
			RequiredPlan:      "Unknown plan",
			RequiredFeature:   model.PaidFeatureAllProfessionalfeatures,
			TrialNotification: true,
		})

		require.Error(t, err)
		require.Equal(t, err.Error(), ": Unable to save notify data.")
		require.Equal(t, http.StatusInternalServerError, statusCode)

	})

	t.Run("error when feature is unknown when notifying on upgrade", func(t *testing.T) {
		th := Setup(t).InitBasic().InitLogin()
		defer th.TearDown()

		statusCode, err := th.Client.NotifyAdmin(&model.NotifyAdminToUpgradeRequest{
			RequiredPlan:    model.LicenseShortSkuProfessional,
			RequiredFeature: "Unknown feature",
		})

		require.Error(t, err)
		require.Equal(t, err.Error(), ": Unable to save notify data.")
		require.Equal(t, http.StatusInternalServerError, statusCode)
	})

	t.Run("error when feature is unknown when notifying to trial", func(t *testing.T) {
		th := Setup(t).InitBasic().InitLogin()
		defer th.TearDown()

		statusCode, err := th.Client.NotifyAdmin(&model.NotifyAdminToUpgradeRequest{
			RequiredPlan:      model.LicenseShortSkuProfessional,
			RequiredFeature:   "Unknown feature",
			TrialNotification: true,
		})

		require.Error(t, err)
		require.Equal(t, err.Error(), ": Unable to save notify data.")
		require.Equal(t, http.StatusInternalServerError, statusCode)
	})

	t.Run("error when user tries to notify again on same feature within the cool off period", func(t *testing.T) {
		th := Setup(t).InitBasic().InitLogin()
		defer th.TearDown()

		statusCode, err := th.Client.NotifyAdmin(&model.NotifyAdminToUpgradeRequest{
			RequiredPlan:    model.LicenseShortSkuProfessional,
			RequiredFeature: model.PaidFeatureAllProfessionalfeatures,
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)

		// second attempt to notify for all professional features
		statusCode, err = th.Client.NotifyAdmin(&model.NotifyAdminToUpgradeRequest{
			RequiredPlan:    model.LicenseShortSkuProfessional,
			RequiredFeature: model.PaidFeatureAllProfessionalfeatures,
		})
		require.Error(t, err)

		require.Equal(t, err.Error(), ": Already notified admin")
		require.Equal(t, http.StatusForbidden, statusCode)
	})

	t.Run("successfully save upgrade notification", func(t *testing.T) {
		th := Setup(t).InitBasic().InitLogin()
		defer th.TearDown()

		statusCode, err := th.Client.NotifyAdmin(&model.NotifyAdminToUpgradeRequest{
			RequiredPlan:    model.LicenseShortSkuProfessional,
			RequiredFeature: model.PaidFeatureAllProfessionalfeatures,
		})

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)
	})
}

func TestTriggerNotifyAdmin(t *testing.T) {
	t.Run("error when EnableTesting is not true", func(t *testing.T) {
		th := Setup(t).InitBasic().InitLogin()
		defer th.TearDown()

		th.App.UpdateConfig(func(cfg *model.Config) { *cfg.ServiceSettings.EnableTesting = false })

		statusCode, err := th.SystemAdminClient.TriggerNotifyAdmin(&model.NotifyAdminToUpgradeRequest{})

		require.Error(t, err)
		require.Equal(t, err.Error(), ": Internal error during cloud api request.")
		require.Equal(t, http.StatusForbidden, statusCode)

	})

	t.Run("error when non admins try to trigger notifications", func(t *testing.T) {
		th := Setup(t).InitBasic().InitLogin()
		defer th.TearDown()

		th.App.UpdateConfig(func(cfg *model.Config) { *cfg.ServiceSettings.EnableTesting = true })

		statusCode, err := th.Client.TriggerNotifyAdmin(&model.NotifyAdminToUpgradeRequest{})

		require.Error(t, err)
		require.Equal(t, err.Error(), ": You do not have the appropriate permissions.")
		require.Equal(t, http.StatusForbidden, statusCode)
	})

	t.Run("happy path", func(t *testing.T) {
		th := Setup(t)
		defer th.TearDown()

		th.App.UpdateConfig(func(cfg *model.Config) { *cfg.ServiceSettings.EnableTesting = true })

		statusCode, err := th.Client.NotifyAdmin(&model.NotifyAdminToUpgradeRequest{
			RequiredPlan:    model.LicenseShortSkuProfessional,
			RequiredFeature: model.PaidFeatureAllProfessionalfeatures,
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)

		statusCode, err = th.SystemAdminClient.TriggerNotifyAdmin(&model.NotifyAdminToUpgradeRequest{})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, statusCode)
	})
}
func Test_validateBusinessEmail(t *testing.T) {
	t.Run("Returns forbidden for non admin executors", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		invalidEmail := model.ValidateBusinessEmailRequest{Email: "invalid@gmail.com"}

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("ValidateBusinessEmail", th.SystemAdminUser.Id, invalidEmail.Email).Return(errors.New("invalid email"))

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		res, err := th.Client.ValidateBusinessEmail(&invalidEmail)
		require.Error(t, err)
		require.Equal(t, http.StatusForbidden, res.StatusCode, "403")
	})

	t.Run("Returns forbidden for invalid business email", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		validBusinessEmail := model.ValidateBusinessEmailRequest{Email: "invalid@slacker.com"}

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("ValidateBusinessEmail", th.SystemAdminUser.Id, validBusinessEmail.Email).Return(errors.New("invalid email"))

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		res, err := th.SystemAdminClient.ValidateBusinessEmail(&validBusinessEmail)
		require.Error(t, err)
		require.Equal(t, http.StatusForbidden, res.StatusCode, "403")
	})

	t.Run("Validate business email for admin", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		validBusinessEmail := model.ValidateBusinessEmailRequest{Email: "valid@mattermost.com"}

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("ValidateBusinessEmail", th.SystemAdminUser.Id, validBusinessEmail.Email).Return(nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		res, err := th.SystemAdminClient.ValidateBusinessEmail(&validBusinessEmail)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode, "200")
	})
}

func Test_validateWorkspaceBusinessEmail(t *testing.T) {
	t.Run("validate the Cloud Customer has used a valid email to create the workspace", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloudCustomerInfo := model.CloudCustomerInfo{
			Email: "valid@mattermost.com",
		}

		cloudCustomer := &model.CloudCustomer{
			CloudCustomerInfo: cloudCustomerInfo,
		}

		cloud.Mock.On("GetCloudCustomer", th.SystemAdminUser.Id).Return(cloudCustomer, nil)
		cloud.Mock.On("ValidateBusinessEmail", th.SystemAdminUser.Id, cloudCustomerInfo.Email).Return(nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		_, err := th.SystemAdminClient.ValidateWorkspaceBusinessEmail()
		require.NoError(t, err)
	})

	t.Run("validate the Cloud Customer has used a invalid email to create the workspace and must validate admin email", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloudCustomerInfo := model.CloudCustomerInfo{
			Email: "invalid@gmail.com",
		}

		cloudCustomer := &model.CloudCustomer{
			CloudCustomerInfo: cloudCustomerInfo,
		}

		cloud.Mock.On("GetCloudCustomer", th.SystemAdminUser.Id).Return(cloudCustomer, nil)

		// first call to validate the cloud customer email
		cloud.Mock.On("ValidateBusinessEmail", th.SystemAdminUser.Id, cloudCustomerInfo.Email).Return(errors.New("invalid email"))

		// second call to validate the user admin email
		cloud.Mock.On("ValidateBusinessEmail", th.SystemAdminUser.Id, th.SystemAdminUser.Email).Return(nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		_, err := th.SystemAdminClient.ValidateWorkspaceBusinessEmail()
		require.NoError(t, err)
	})
}

func TestGetCloudProducts(t *testing.T) {
	cloudProducts := []*model.Product{
		{
			ID:                "prod_test1",
			Name:              "name",
			Description:       "description",
			PricePerSeat:      10,
			SKU:               "sku",
			PriceID:           "price_id",
			Family:            "family",
			RecurringInterval: "recurring_interval",
			BillingScheme:     "billing_scheme",
		},
		{
			ID:                "prod_test2",
			Name:              "name2",
			Description:       "description2",
			PricePerSeat:      100,
			SKU:               "sku2",
			PriceID:           "price_id2",
			Family:            "family2",
			RecurringInterval: "recurring_interval2",
			BillingScheme:     "billing_scheme2",
		},
		{
			ID:                "prod_test3",
			Name:              "name3",
			Description:       "description3",
			PricePerSeat:      1000,
			SKU:               "sku3",
			PriceID:           "price_id3",
			Family:            "family3",
			RecurringInterval: "recurring_interval3",
			BillingScheme:     "billing_scheme3",
		},
	}

	sanitizedProducts := []*model.Product{
		{
			ID:           "prod_test1",
			Name:         "name",
			PricePerSeat: 10,
			SKU:          "sku",
		},
		{
			ID:           "prod_test2",
			Name:         "name2",
			PricePerSeat: 100,
			SKU:          "sku2",
		},
		{
			ID:           "prod_test3",
			Name:         "name3",
			PricePerSeat: 1000,
			SKU:          "sku3",
		},
	}
	t.Run("get products for admins", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.SystemAdminUser.Email, th.SystemAdminUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}
		cloud.Mock.On("GetCloudProducts", mock.Anything, mock.Anything).Return(cloudProducts, nil)
		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		returnedProducts, r, err := th.Client.GetCloudProducts()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, r.StatusCode, "Status OK")
		require.Equal(t, returnedProducts, cloudProducts)
	})

	t.Run("get products for non admins", func(t *testing.T) {
		th := Setup(t).InitBasic()
		defer th.TearDown()

		th.Client.Login(th.BasicUser.Email, th.BasicUser.Password)

		th.App.Srv().SetLicense(model.NewTestLicense("cloud"))

		cloud := mocks.CloudInterface{}

		cloud.Mock.On("GetCloudProducts", mock.Anything, mock.Anything).Return(cloudProducts, nil)

		cloudImpl := th.App.Srv().Cloud
		defer func() {
			th.App.Srv().Cloud = cloudImpl
		}()
		th.App.Srv().Cloud = &cloud

		returnedProducts, r, err := th.Client.GetCloudProducts()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, r.StatusCode, "Status OK")
		require.Equal(t, returnedProducts, sanitizedProducts)

		// make a more explicit check
		require.Equal(t, returnedProducts[0].ID, "prod_test1")
		require.Equal(t, returnedProducts[0].Name, "name")
		require.Equal(t, returnedProducts[0].SKU, "sku")
		require.Equal(t, returnedProducts[0].PricePerSeat, float64(10))
		require.Equal(t, returnedProducts[0].Description, "")
		require.Equal(t, returnedProducts[0].PriceID, "")
		require.Equal(t, returnedProducts[0].Family, model.SubscriptionFamily(""))
		require.Equal(t, returnedProducts[0].RecurringInterval, model.RecurringInterval(""))
		require.Equal(t, returnedProducts[0].BillingScheme, model.BillingScheme(""))

		require.Equal(t, returnedProducts[1].ID, "prod_test2")
		require.Equal(t, returnedProducts[1].Name, "name2")
		require.Equal(t, returnedProducts[1].SKU, "sku2")
		require.Equal(t, returnedProducts[1].PricePerSeat, float64(100))
		require.Equal(t, returnedProducts[1].Description, "")
		require.Equal(t, returnedProducts[1].PriceID, "")
		require.Equal(t, returnedProducts[1].Family, model.SubscriptionFamily(""))
		require.Equal(t, returnedProducts[1].RecurringInterval, model.RecurringInterval(""))
		require.Equal(t, returnedProducts[1].BillingScheme, model.BillingScheme(""))

		require.Equal(t, returnedProducts[2].ID, "prod_test3")
		require.Equal(t, returnedProducts[2].Name, "name3")
		require.Equal(t, returnedProducts[2].SKU, "sku3")
		require.Equal(t, returnedProducts[2].PricePerSeat, float64(1000))
		require.Equal(t, returnedProducts[2].Description, "")
		require.Equal(t, returnedProducts[2].PriceID, "")
		require.Equal(t, returnedProducts[2].Family, model.SubscriptionFamily(""))
		require.Equal(t, returnedProducts[2].RecurringInterval, model.RecurringInterval(""))
		require.Equal(t, returnedProducts[2].BillingScheme, model.BillingScheme(""))
	})
}

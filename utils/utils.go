// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package utils

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/v5/model"
)

func StringInSlice(a string, slice []string) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}

// RemoveStringFromSlice removes the first occurrence of a from slice.
func RemoveStringFromSlice(a string, slice []string) []string {
	for i, str := range slice {
		if str == a {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// RemoveStringsFromSlice removes all occurrences of strings from slice.
func RemoveStringsFromSlice(slice []string, strings ...string) []string {
	newSlice := []string{}

	for _, item := range slice {
		if !StringInSlice(item, strings) {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}

func StringArrayIntersection(arr1, arr2 []string) []string {
	arrMap := map[string]bool{}
	result := []string{}

	for _, value := range arr1 {
		arrMap[value] = true
	}

	for _, value := range arr2 {
		if arrMap[value] {
			result = append(result, value)
		}
	}

	return result
}

func RemoveDuplicatesFromStringArray(arr []string) []string {
	result := make([]string, 0, len(arr))
	seen := make(map[string]bool)

	for _, item := range arr {
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}

func StringSliceDiff(a, b []string) []string {
	m := make(map[string]bool)
	result := []string{}

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if !m[item] {
			result = append(result, item)
		}
	}
	return result
}

func GetIpAddress(r *http.Request, trustedProxyIPHeader []string) string {
	address := ""

	for _, proxyHeader := range trustedProxyIPHeader {
		header := r.Header.Get(proxyHeader)
		if len(header) > 0 {
			addresses := strings.Fields(header)
			if len(addresses) > 0 {
				address = strings.TrimRight(addresses[0], ",")
			}
		}

		if len(address) > 0 {
			return address
		}
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return address
}

func GetHostnameFromSiteURL(siteURL string) string {
	u, err := url.Parse(siteURL)
	if err != nil {
		return ""
	}

	return u.Hostname()
}

type RequestCache struct {
	Data []byte
	Date string
	Key  string
}

// Fetch JSON data from the notices server
// if skip is passed, does a fetch without touching the cache
func GetUrlWithCache(url string, cache *RequestCache, skip bool) ([]byte, error) {
	// Build a GET Request, including optional If-None-Match header.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		cache.Data = nil
		return nil, err
	}
	if !skip && cache.Data != nil {
		req.Header.Add("If-None-Match", cache.Key)
		req.Header.Add("If-Modified-Since", cache.Date)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		cache.Data = nil
		return nil, err
	}
	defer resp.Body.Close()
	// No change from latest known Etag?
	if resp.StatusCode == http.StatusNotModified {
		return cache.Data, nil
	}

	if resp.StatusCode != 200 {
		cache.Data = nil
		return nil, errors.Errorf("Fetching notices failed with status code %d", resp.StatusCode)
	}

	cache.Data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		cache.Data = nil
		return nil, err
	}

	// If etags headers are missing, ignore.
	cache.Key = resp.Header.Get("ETag")
	cache.Date = resp.Header.Get("Date")
	return cache.Data, err
}

// Append cookies to passed baseUrl into query params
func BuildUrlQueryStringFromCookies(baseUrl string, r *http.Request) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return ""
	}
	for _, cookie := range r.Cookies() {
		q.Add(cookie.Name, cookie.Value)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// Append tokens to passed baseUrl as query params
func BuildUrlQueryStringWithTokenInfo(baseUrl string, token string, csrf string) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return ""
	}
	q.Add(model.SESSION_COOKIE_TOKEN, token)
	q.Add(model.SESSION_COOKIE_CSRF, csrf)
	u.RawQuery = q.Encode()
	return u.String()
}

// Validates RedirectURL passed during OAuth or SAML
func IsValidWebAuthRedirectURL(config *model.Config, redirectUrl string) bool {
	if config.ServiceSettings.SiteURL != nil {
		siteUrl := *config.ServiceSettings.SiteURL
		return strings.Index(strings.ToLower(redirectUrl), strings.ToLower(siteUrl)) == 0
	}
	return false
}

// Validates Mobile Custom URL Scheme passed during OAuth or SAML
func IsValidMobileAuthRedirectURL(config *model.Config, redirectUrl string) bool {
	if config.NativeAppSettings.AppCustomUrlScheme != nil {
		urlScheme := *config.NativeAppSettings.AppCustomUrlScheme
		return strings.Index(strings.ToLower(redirectUrl), strings.ToLower(urlScheme)) == 0
	}
	return false
}

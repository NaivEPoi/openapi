//go:build !debug
// +build !debug

/*
 * Nausf_SoRProtection Service
 *
 * AUSF SoR Protection Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nausf_SoRProtection

// APIClient manages communication with the Nausf_SoRProtection Service API v1.0.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg	*Configuration
	common	service	// Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	DefaultApi	*DefaultApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.DefaultApi = (*DefaultApiService)(&c.common)

	return c
}

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	RawBody		[]byte
	ErrorStatus	string
	ErrorModel	interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.ErrorStatus
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.RawBody
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.ErrorModel
}

/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PostSmContextsErrorResponse struct {
	JsonData		*SmContextCreateError	`json:"jsonData,omitempty" multipart:"contentType:application/json"`
	BinaryDataN1SmMessage	[]byte			`json:"binaryDataN1SmMessage,omitempty" multipart:"contentType:application/vnd.3gpp.5gnas,ref:JsonData.N1SmMsg.ContentId"`
}

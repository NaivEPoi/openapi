/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type StatusInfo struct {
	ResourceStatus	ResourceStatus	`json:"resourceStatus"`
	Cause		Cause		`json:"cause,omitempty"`
}

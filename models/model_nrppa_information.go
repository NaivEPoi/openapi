/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NrppaInformation struct {
	NfId		string		`json:"nfId"`
	NrppaPdu	*N2InfoContent	`json:"nrppaPdu"`
}

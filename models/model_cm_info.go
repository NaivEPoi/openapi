/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type CmInfo struct {
	CmState		CmState		`json:"cmState" bson:"cmState"`
	AccessType	AccessType	`json:"accessType" bson:"accessType"`
}

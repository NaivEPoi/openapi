/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ExposureDataChangeNotification struct {
	UeId				string				`json:"ueId,omitempty" bson:"ueId"`
	AccessAndMobilityData		*AccessAndMobilityData		`json:"accessAndMobilityData,omitempty" bson:"accessAndMobilityData"`
	PduSessionManagementData	[]PduSessionManagementData	`json:"pduSessionManagementData,omitempty" bson:"pduSessionManagementData"`
}

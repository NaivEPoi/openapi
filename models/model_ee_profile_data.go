/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type EeProfileData struct {
	RestrictedEventTypes	[]EventType	`json:"restrictedEventTypes,omitempty" bson:"restrictedEventTypes"`
	SupportedFeatures	string		`json:"supportedFeatures,omitempty" bson:"supportedFeatures"`
}

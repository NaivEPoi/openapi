/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ExposureDataSubscription struct {
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationUri		string		`json:"notificationUri" bson:"notificationUri"`
	MonitoredResourceUris	[]string	`json:"monitoredResourceUris" bson:"monitoredResourceUris"`
	SupportedFeatures	string		`json:"supportedFeatures,omitempty" bson:"supportedFeatures"`
}

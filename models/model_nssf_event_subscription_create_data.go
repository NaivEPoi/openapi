/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type NssfEventSubscriptionCreateData struct {
	NfNssaiAvailabilityUri	string	`json:"nfNssaiAvailabilityUri" bson:"nfNssaiAvailabilityUri" yaml:"nfNssaiAvailabilityUri"`

	TaiList	[]Tai	`json:"taiList" bson:"taiList" yaml:"taiList"`

	Event	NssfEventType	`json:"event" bson:"event" yaml:"event"`

	Expiry	*time.Time	`json:"expiry,omitempty" bson:"expiry" yaml:"expiry"`
}

/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type NrLocation struct {
	Tai				*Tai			`json:"tai" yaml:"tai" bson:"tai"`
	Ncgi				*Ncgi			`json:"ncgi" yaml:"ncgi" bson:"ncgi"`
	AgeOfLocationInformation	int32			`json:"ageOfLocationInformation,omitempty" yaml:"ageOfLocationInformation" bson:"ageOfLocationInformation"`
	UeLocationTimestamp		*time.Time		`json:"ueLocationTimestamp,omitempty" yaml:"ueLocationTimestamp" bson:"ueLocationTimestamp"`
	GeographicalInformation		string			`json:"geographicalInformation,omitempty" yaml:"geographicalInformation" bson:"geographicalInformation"`
	GeodeticInformation		string			`json:"geodeticInformation,omitempty" yaml:"geodeticInformation" bson:"geodeticInformation"`
	GlobalGnbId			*GlobalRanNodeId	`json:"globalGnbId,omitempty" yaml:"globalGnbId" bson:"globalGnbId"`
}

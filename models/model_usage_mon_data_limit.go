/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Contains usage monitoring control data for a subscriber.
type UsageMonDataLimit struct {
	LimitId		string				`json:"limitId" bson:"limitId"`
	Scopes		map[string]UsageMonDataScope	`json:"scopes,omitempty" bson:"scopes"`
	UmLevel		UsageMonLevel			`json:"umLevel,omitempty" bson:"umLevel"`
	StartDate	*time.Time			`json:"startDate,omitempty" bson:"startDate"`
	EndDate		*time.Time			`json:"endDate,omitempty" bson:"endDate"`
	UsageLimit	*UsageThreshold			`json:"usageLimit,omitempty" bson:"usageLimit"`
	ResetPeriod	*time.Time			`json:"resetPeriod,omitempty" bson:"resetPeriod"`
}

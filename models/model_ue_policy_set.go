/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the UE policy data for a given subscriber.
type UePolicySet struct {
	SubscCats		[]string			`json:"subscCats,omitempty" bson:"subscCats"`
	UePolicySections	map[string]UePolicySection	`json:"uePolicySections,omitempty" bson:"uePolicySections"`
	Upsis			[]string			`json:"upsis,omitempty" bson:"upsis"`
}

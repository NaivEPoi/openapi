/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PduSessionTypes struct {
	DefaultSessionType	PduSessionType		`json:"defaultSessionType" yaml:"defaultSessionType" bson:"defaultSessionType" mapstructure:"DefaultSessionType"`
	AllowedSessionTypes	[]PduSessionType	`json:"allowedSessionTypes,omitempty" yaml:"allowedSessionTypes" bson:"allowedSessionTypes" mapstructure:"AllowedSessionTypes"`
}

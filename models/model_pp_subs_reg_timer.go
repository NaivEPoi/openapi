/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PpSubsRegTimer struct {
	SubsRegTimer	int32	`json:"subsRegTimer" yaml:"subsRegTimer" bson:"subsRegTimer" mapstructure:"SubsRegTimer"`
	AfInstanceId	string	`json:"afInstanceId" yaml:"afInstanceId" bson:"afInstanceId" mapstructure:"AfInstanceId"`
	ReferenceId	int32	`json:"referenceId" yaml:"referenceId" bson:"referenceId" mapstructure:"ReferenceId"`
}

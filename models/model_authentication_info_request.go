/*
 * NudmUEAU
 *
 * UDM UE Authentication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AuthenticationInfoRequest struct {
	SupportedFeatures	string			`json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
	ServingNetworkName	string			`json:"servingNetworkName" yaml:"servingNetworkName" bson:"servingNetworkName" mapstructure:"ServingNetworkName"`
	ResynchronizationInfo	*ResynchronizationInfo	`json:"resynchronizationInfo,omitempty" yaml:"resynchronizationInfo" bson:"resynchronizationInfo" mapstructure:"ResynchronizationInfo"`
	AusfInstanceId		string			`json:"ausfInstanceId" yaml:"ausfInstanceId" bson:"ausfInstanceId" mapstructure:"AusfInstanceId"`
}

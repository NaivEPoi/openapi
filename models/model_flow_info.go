/*
 * 3gpp-traffic-influence
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type FlowInfo struct {
	// Indicates the IP flow.
	FlowId	int32	`json:"flowId" yaml:"flowId" bson:"flowId" mapstructure:"FlowId"`
	// Indicates the packet filters of the IP flow. Refer to subclause 5.3.8 of 3GPP TS 29.214 for encoding. It shall contain UL and/or DL IP flow description.
	FlowDescriptions	[]string	`json:"flowDescriptions,omitempty" yaml:"flowDescriptions" bson:"flowDescriptions" mapstructure:"FlowDescriptions"`
}

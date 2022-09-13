/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type QosFlowProfile struct {
	Var5qi			int32			`json:"5qi"`
	NonDynamic5Qi		*NonDynamic5Qi		`json:"nonDynamic5Qi,omitempty"`
	Dynamic5Qi		*Dynamic5Qi		`json:"dynamic5Qi,omitempty"`
	Arp			*Arp			`json:"arp,omitempty"`
	GbrQosFlowInfo		*GbrQosFlowInformation	`json:"gbrQosFlowInfo,omitempty"`
	Rqa			ReflectiveQoSAttribute	`json:"rqa,omitempty"`
	AdditionalQosFlowInfo	AdditionalQosFlowInfo	`json:"additionalQosFlowInfo,omitempty"`
}

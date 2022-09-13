/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SessionRuleFailureCode string

// List of SessionRuleFailureCode
const (
	SessionRuleFailureCode_NF_MAL		SessionRuleFailureCode	= "NF_MAL"
	SessionRuleFailureCode_RES_LIM		SessionRuleFailureCode	= "RES_LIM"
	SessionRuleFailureCode_UNSUCC_QOS_VAL	SessionRuleFailureCode	= "UNSUCC_QOS_VAL"
	SessionRuleFailureCode_UE_STA_SUSP	SessionRuleFailureCode	= "UE_STA_SUSP"
)

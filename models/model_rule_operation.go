/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RuleOperation string

// List of RuleOperation
const (
	RuleOperation_CREATE_PCC_RULE					RuleOperation	= "CREATE_PCC_RULE"
	RuleOperation_DELETE_PCC_RULE					RuleOperation	= "DELETE_PCC_RULE"
	RuleOperation_MODIFY_PCC_RULE_AND_ADD_PACKET_FILTERS		RuleOperation	= "MODIFY_PCC_RULE_AND_ADD_PACKET_FILTERS"
	RuleOperation_MODIFY_PCC_RULE_AND_REPLACE_PACKET_FILTERS	RuleOperation	= "MODIFY_PCC_RULE_AND_REPLACE_PACKET_FILTERS"
	RuleOperation_MODIFY_PCC_RULE_AND_DELETE_PACKET_FILTERS		RuleOperation	= "MODIFY_PCC_RULE_AND_DELETE_PACKET_FILTERS"
	RuleOperation_MODIFY_PCC_RULE_WITHOUT_MODIFY_PACKET_FILTERS	RuleOperation	= "MODIFY_PCC_RULE_WITHOUT_MODIFY_PACKET_FILTERS"
)

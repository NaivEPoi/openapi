/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents an Individual Application Session Context resource.
type AppSessionContext struct {
	AscReqData	*AppSessionContextReqData	`json:"ascReqData,omitempty" yaml:"ascReqData" bson:"ascReqData" mapstructure:"AscReqData"`
	AscRespData	*AppSessionContextRespData	`json:"ascRespData,omitempty" yaml:"ascRespData" bson:"ascRespData" mapstructure:"AscRespData"`
	EvsNotif	*EventsNotification		`json:"evsNotif,omitempty" yaml:"evsNotif" bson:"evsNotif" mapstructure:"EvsNotif"`
}

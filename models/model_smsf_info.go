/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SmsfInfo struct {
	SmsfInstanceId	string	`json:"smsfInstanceId" yaml:"smsfInstanceId" bson:"smsfInstanceId" mapstructure:"SmsfInstanceId"`
	PlmnId		*PlmnId	`json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
}

/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type AcknowledgeInfo struct {
	SorMacIue		string		`json:"sorMacIue,omitempty" yaml:"sorMacIue" bson:"sorMacIue" mapstructure:"SorMacIue"`
	UpuMacIue		string		`json:"upuMacIue,omitempty" yaml:"upuMacIue" bson:"upuMacIue" mapstructure:"UpuMacIue"`
	SecuredPacket		string		`json:"securedPacket,omitempty" yaml:"securedPacket" bson:"securedPacket" mapstructure:"SecuredPacket"`
	ProvisioningTime	*time.Time	`json:"provisioningTime" yaml:"provisioningTime" bson:"provisioningTime" mapstructure:"ProvisioningTime"`
}

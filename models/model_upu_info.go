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

type UpuInfo struct {
	UpuDataList		[]UpuData	`json:"upuDataList" yaml:"upuDataList" bson:"upuDataList" mapstructure:"UpuDataList"`
	UpuRegInd		bool		`json:"upuRegInd" yaml:"upuRegInd" bson:"upuRegInd" mapstructure:"UpuRegInd"`
	UpuAckInd		bool		`json:"upuAckInd" yaml:"upuAckInd" bson:"upuAckInd" mapstructure:"UpuAckInd"`
	UpuMacIausf		string		`json:"upuMacIausf,omitempty" yaml:"upuMacIausf" bson:"upuMacIausf" mapstructure:"UpuMacIausf"`
	CounterUpu		string		`json:"counterUpu,omitempty" yaml:"counterUpu" bson:"counterUpu" mapstructure:"CounterUpu"`
	ProvisioningTime	*time.Time	`json:"provisioningTime" yaml:"provisioningTime" bson:"provisioningTime" mapstructure:"ProvisioningTime"`
}

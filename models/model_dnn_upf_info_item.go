/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type DnnUpfInfoItem struct {
	Dnn		string			`json:"dnn" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
	DnaiList	[]string		`json:"dnaiList,omitempty" yaml:"dnaiList" bson:"dnaiList" mapstructure:"DnaiList"`
	PduSessionTypes	[]PduSessionType	`json:"pduSessionTypes,omitempty" yaml:"pduSessionTypes" bson:"pduSessionTypes" mapstructure:"PduSessionTypes"`
}

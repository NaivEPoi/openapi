/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Ecgi struct {
	PlmnId		*PlmnId	`json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	EutraCellId	string	`json:"eutraCellId" yaml:"eutraCellId" bson:"eutraCellId"`
}

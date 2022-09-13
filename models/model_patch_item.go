/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PatchItem struct {
	Op	PatchOperation	`json:"op" yaml:"op" bson:"op" mapstructure:"Op"`
	Path	string		`json:"path" yaml:"path" bson:"path" mapstructure:"Path"`
	From	string		`json:"from,omitempty" yaml:"from" bson:"from" mapstructure:"From"`
	Value	interface{}	`json:"value,omitempty" yaml:"value" bson:"value" mapstructure:"Value"`
}

/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type BackupAmfInfo struct {
	BackupAmf	string	`json:"backupAmf" yaml:"backupAmf" bson:"backupAmf" mapstructure:"BackupAmf"`
	GuamiList	[]Guami	`json:"guamiList,omitempty" yaml:"guamiList" bson:"guamiList" mapstructure:"GuamiList"`
}

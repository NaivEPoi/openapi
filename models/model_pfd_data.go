/*
 * 3gpp-pfd-management
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PfdData struct {
	// Each element uniquely external application identifier
	ExternalAppId	string	`json:"externalAppId" yaml:"externalAppId" bson:"externalAppId" mapstructure:"ExternalAppId"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self	string	`json:"self,omitempty" yaml:"self" bson:"self" mapstructure:"Self"`
	// Contains the PFDs of the external application identifier. Each PFD is identified in the map via a key containing the PFD identifier.
	Pfds	map[string]Pfd	`json:"pfds" yaml:"pfds" bson:"pfds" mapstructure:"Pfds"`
	// Unsigned integer identifying a period of time in units of seconds with \"nullable=true\" property.
	AllowedDelay	int32	`json:"allowedDelay,omitempty" yaml:"allowedDelay" bson:"allowedDelay" mapstructure:"AllowedDelay"`
	// Unsigned integer identifying a period of time in units of seconds with \"readOnly=true\" property.
	CachingTime	int32	`json:"cachingTime,omitempty" yaml:"cachingTime" bson:"cachingTime" mapstructure:"CachingTime"`
}

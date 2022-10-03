/*
 * 3gpp-traffic-influence
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type WebsockNotifConfig struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	WebsocketUri	string	`json:"websocketUri,omitempty" yaml:"websocketUri" bson:"websocketUri" mapstructure:"WebsocketUri"`
	// Set by the SCS/AS to indicate that the Websocket delivery is requested.
	RequestWebsocketUri	bool	`json:"requestWebsocketUri,omitempty" yaml:"requestWebsocketUri" bson:"requestWebsocketUri" mapstructure:"RequestWebsocketUri"`
}

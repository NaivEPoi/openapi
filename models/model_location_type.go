/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type LocationType string

// List of LocationType
const (
	LocationType_CURRENT_LOCATION			LocationType	= "CURRENT_LOCATION"
	LocationType_CURRENT_OR_LAST_KNOWN_LOCATION	LocationType	= "CURRENT_OR_LAST_KNOWN_LOCATION"
	LocationType_INITIAL_LOCATION			LocationType	= "INITIAL_LOCATION"
)

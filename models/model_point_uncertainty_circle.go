/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PointUncertaintyCircle struct {
	Shape		SupportedGadShapes		`json:"shape" yaml:"shape" bson:"shape"`
	Point		*GeographicalCoordinates	`json:"point" yaml:"point" bson:"point"`
	Uncertainty	float32				`json:"uncertainty" yaml:"uncertainty" bson:"uncertainty"`
}

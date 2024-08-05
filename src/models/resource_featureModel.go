/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

type ResourceFeature struct {

	// id of the feature.
	Id string `json:"id,omitempty" bson:"id,omitempty"`

	// A boolean. True if this is a feature group. Default is false
	IsBundle bool `json:"isBundle,omitempty" bson:"isBundle,omitempty"`

	// A boolean. True if this feature is enabled. Default is true
	IsEnabled bool `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`

	// A string used to give a name to the feature
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// free-text description of the feature
	Description string `json:"description,omitempty" bson:"description,omitempty"`

	ResourceFeatureCharacteristic []Characteristic `json:"resourceFeatureCharacteristic,omitempty" bson:"resourceFeatureCharacteristic,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`
}

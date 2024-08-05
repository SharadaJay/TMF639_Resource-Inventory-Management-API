/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// Describes a given characteristic of an object or entity through a name/value pair.
type Characteristic struct {

	// Unique identifier of the characteristic
	Id string `json:"id,omitempty" bson:"id,omitempty"`

	// Name of the characteristic
	Name string `json:"name" bson:"name"`

	// Data type of the value of the characteristic
	ValueType string `json:"valueType,omitempty" bson:"valueType,omitempty"`

	CharacteristicRelationship []CharacteristicRelationship `json:"characteristicRelationship,omitempty" bson:"characteristicRelationship,omitempty"`

	// The value of the characteristic Edited by spj
	// Value *Any `json:"value"` Commented by spj
	Value string `json:"value" bson:"value"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`
}

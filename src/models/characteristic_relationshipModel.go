/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// Another Characteristic that is related to the current Characteristic;
type CharacteristicRelationship struct {

	// Unique identifier of the characteristic
	Id string `json:"id,omitempty" bson:"id,omitempty"`

	// The type of relationship
	RelationshipType string `json:"relationshipType,omitempty" bson:"relationshipType,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`
}

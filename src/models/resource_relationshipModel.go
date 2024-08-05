/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

type ResourceRelationship struct {

	// id of the related resource.
	Id string `json:"id,omitempty" bson:"id,omitempty"`

	// href of the related resource.
	Href string `json:"href,omitempty" bson:"href,omitempty"`

	RelationshipType string `json:"relationshipType,omitempty" bson:"relationshipType,omitempty"`

	ResourceRelationshipCharacteristic []Characteristic `json:"resourceRelationshipCharacteristic,omitempty" bson:"resourceRelationshipCharacteristic,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`
}

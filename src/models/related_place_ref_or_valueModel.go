/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// Related Entity reference. A related place defines a place described by reference or by value linked to a specific entity. The polymorphic attributes @type, @schemaLocation & @referredType are related to the place entity and not the RelatedPlaceRefOrValue class itself
type RelatedPlaceRefOrValue struct {

	// Unique identifier of the place
	Id string `json:"id" bson:"id"`

	// Unique reference of the place
	Href string `json:"href" bson:"href"`

	// A user-friendly name for the place, such as \"Paris Store\", \"London Store\", \"Main Home\"
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	Role string `json:"role" bson:"role"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty" bson:"@referredType,omitempty"`
}

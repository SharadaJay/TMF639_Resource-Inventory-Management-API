/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// Related Entity reference. A related party defines party or party role linked to a specific entity.
type RelatedParty struct {

	// Unique identifier of a related entity.
	Id string `json:"id" bson:"id"`

	// Reference of the related entity.
	Href string `json:"href,omitempty" bson:"href,omitempty"`

	// Name of the related entity.
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// Role played by the related party
	Role string `json:"role,omitempty" bson:"role,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType" bson:"@referredType"`
}

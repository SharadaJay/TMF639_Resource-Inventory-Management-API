/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// Resource Specification reference: The ResourceSpecification is required to realize a ProductSpecification.
type ResourceSpecificationRef struct {

	// Unique identifier of the resource specification
	Id string `json:"id" bson:"id"`

	// Reference of the resource specification
	Href string `json:"href" bson:"href"`

	// Name of the requiredResourceSpecification
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// Resource specification version
	Version string `json:"version,omitempty" bson:"version,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty" bson:"@referredType,omitempty"`
}

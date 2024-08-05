/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

import (
	"time"
)

// Extra information about a given entity
type Note struct {

	// Identifier of the note within its containing entity (may or may not be globally unique, depending on provider implementation)
	Id string `json:"id,omitempty" bson:"id,omitempty"`

	// Author of the note
	Author string `json:"author,omitempty" bson:"author,omitempty"`

	// Date of the note
	Date time.Time `json:"date,omitempty" bson:"date,omitempty"`

	// Text of the note
	Text string `json:"text,omitempty" bson:"text,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`
}

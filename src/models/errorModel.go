/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// Used when an API throws an Error, typically with a HTTP error response-code (3xx, 4xx, 5xx)
type ModelError struct {

	// Application relevant detail, defined in the API or a common list.
	Code string `json:"code" bson:"code"`

	// Explanation of the reason for the error which can be shown to a client user.
	Reason string `json:"reason" bson:"reason"`

	// More details and corrective actions related to the error which can be shown to a client user.
	Message string `json:"message,omitempty" bson:"message,omitempty"`

	// HTTP Error code extension
	Status string `json:"status,omitempty" bson:"status,omitempty"`

	// URI of documentation describing the error.
	ReferenceError string `json:"referenceError,omitempty" bson:"referenceError,omitempty"`

	// When sub-classing, this defines the super-class.
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name.
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`
}

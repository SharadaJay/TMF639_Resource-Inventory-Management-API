/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// An attachment by value or by reference. For AttachmentRefOrValue, the attribute type,schemaLocation and referredType are related to the contained entity and not to AttchmentRefOrValue itself
type AttachmentRefOrValue struct {

	// Unique identifier for this particular attachment
	Id string `json:"id,omitempty" bson:"id,omitempty"`

	// URI for this Attachment
	Href string `json:"href,omitempty" bson:"href,omitempty"`

	// Attachment type such as video, picture
	AttachmentType string `json:"attachmentType,omitempty" bson:"attachmentType,omitempty"`

	// The actual contents of the attachment object, if embedded, encoded as base64
	Content string `json:"content,omitempty" bson:"content,omitempty"`

	// A narrative text describing the content of the attachment
	Description string `json:"description,omitempty" bson:"description,omitempty"`

	// Attachment mime type such as extension file for video, picture and document
	MimeType string `json:"mimeType,omitempty" bson:"mimeType,omitempty"`

	// The name of the attachment
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// Uniform Resource Locator, is a web page address (a subset of URI)
	Url string `json:"url,omitempty" bson:"url,omitempty"`

	// The size of the attachment.
	Size *Quantity `json:"size,omitempty" bson:"size,omitempty"`

	// The period of time for which the attachment is valid
	ValidFor *TimePeriod `json:"validFor,omitempty" bson:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty" bson:"@referredType,omitempty"`
}

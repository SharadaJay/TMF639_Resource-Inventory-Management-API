/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

import (
	"time"
)

// Resource is an abstract entity that describes the common set of attributes shared by all concrete resources (e.g. TPE, EQUIPMENT) in the inventory.
type Resource struct {

	// Identifier of an instance of the resource. Required to be unique within the resource type.  Used in URIs as the identifier for specific instances of a type.
	Id string `json:"id" bson:"id"`

	// The URI for the object itself.
	Href string `json:"href" bson:"href"`

	// Category of the concrete resource. e.g Gold, Silver for MSISDN concrete resource
	Category string `json:"category,omitempty" bson:"category,omitempty"`

	// Value (Added by SPJ)
	Value string `json:"value,omitempty" bson:"value,omitempty"`

	// free-text description of the resource
	Description string `json:"description,omitempty" bson:"description,omitempty"`

	// A date time( DateTime). The date till the resource is operating
	EndOperatingDate *time.Time `json:"endOperatingDate,omitempty" bson:"endOperatingDate,omitempty"`

	// A string used to give a name to the resource
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// A string used to give the relevant Account Id to the resource
	AccountId string `json:"accountId,omitempty" bson:"accountId,omitempty"`

	// A string used to give athe relevant Cloude Resource Id to the resource
	CloudResourceId string `json:"cloudResourceId,omitempty" bson:"cloudResourceId,omitempty"`

	// A field that identifies the specific version of an instance of a resource.
	ResourceVersion string `json:"resourceVersion,omitempty" bson:"resourceVersion,omitempty"`

	// A date time( DateTime). The date from which the resource is operating
	StartOperatingDate *time.Time `json:"startOperatingDate,omitempty" bson:"startOperatingDate,omitempty"`

	// Tracks the lifecycle status of the resource, such as planning, installing, opereating, retiring and so on.
	AdministrativeState *ResourceAdministrativeStateType `json:"administrativeState,omitempty" bson:"administrativeState,omitempty"`

	Attachment []AttachmentRefOrValue `json:"attachment,omitempty" bson:"attachment,omitempty"`

	Note []Note `json:"note,omitempty" bson:"note,omitempty"`

	// Tracks the lifecycle status of the resource, such as planning, installing, opereating, retiring and so on.
	OperationalState *ResourceOperationalStateType `json:"operationalState,omitempty" bson:"operationalState,omitempty"`

	Place *RelatedPlaceRefOrValue `json:"place,omitempty" bson:"place,omitempty"`

	RelatedParty []RelatedParty `json:"relatedParty,omitempty" bson:"relatedParty,omitempty"`

	ResourceCharacteristic []Characteristic `json:"resourceCharacteristic,omitempty" bson:"resourceCharacteristic,omitempty"`

	ResourceRelationship []ResourceRelationship `json:"resourceRelationship,omitempty" bson:"resourceRelationship,omitempty"`

	ResourceSpecification *ResourceSpecificationRef `json:"resourceSpecification,omitempty" bson:"resourceSpecification,omitempty"`

	//Added Resource feature (SPJ)
	ResourceFeature []ResourceFeature `json:"resourceFeature,omitempty" bson:"resourceFeature,omitempty"`

	// Tracks the lifecycle status of the resource, such as planning, installing, opereating, retiring and so on.
	ResourceStatus *ResourceStatusType `json:"resourceStatus,omitempty" bson:"resourceStatus,omitempty"`

	// Tracks the lifecycle status of the resource, such as planning, installing, opereating, retiring and so on.
	UsageState *ResourceUsageStateType `json:"usageState,omitempty" bson:"usageState,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a json-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty" bson:"@type,omitempty"`
}

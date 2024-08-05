/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// ResourceStatusType : ResourceStatusType enumerations
type ResourceStatusType string

// List of ResourceStatusType
const (
	STANDBY   ResourceStatusType = "standby"
	ALARM     ResourceStatusType = "alarm"
	AVAILABLE ResourceStatusType = "available"
	RESERVED  ResourceStatusType = "reserved"
	UNKNOWN   ResourceStatusType = "unknown"
	SUSPENDED ResourceStatusType = "suspended"
)

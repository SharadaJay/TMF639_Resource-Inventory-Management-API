/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// ResourceOperationalStateType : ResourceOperationalStateType enumerations
type ResourceOperationalStateType string

// List of ResourceOperationalStateType
const (
	ENABLE  ResourceOperationalStateType = "enable"
	DISABLE ResourceOperationalStateType = "disable"
)

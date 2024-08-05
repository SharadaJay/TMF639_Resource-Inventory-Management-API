/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// ResourceAdministrativeStateType : ResourceAdministrativeStateType enumerations
type ResourceAdministrativeStateType string

// List of ResourceAdministrativeStateType
const (
	LOCKED   ResourceAdministrativeStateType = "locked"
	UNLOCKED ResourceAdministrativeStateType = "unlocked"
	SHUTDOWN ResourceAdministrativeStateType = "shutdown"
)

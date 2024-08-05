/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// ResourceUsageStateType : ResourceUsageStateType enumerations
type ResourceUsageStateType string

// List of ResourceUsageStateType
const (
	IDLE   ResourceUsageStateType = "idle"
	ACTIVE ResourceUsageStateType = "active"
	BUSY   ResourceUsageStateType = "busy"
)

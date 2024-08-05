/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

// An amount in a given unit
type Quantity struct {

	// Numeric value in a given unit
	Amount float32 `json:"amount,omitempty" bson:"amount,omitempty"`

	// Unit
	Units string `json:"units,omitempty" bson:"units,omitempty"`
}

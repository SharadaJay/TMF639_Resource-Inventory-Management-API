/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package models

import (
	"time"
)

// A period of time, either as a deadline (endDateTime only) a startDateTime only, or both
type TimePeriod struct {

	// End of the time period, using IETC-RFC-3339 format
	EndDateTime *time.Time `json:"endDateTime,omitempty" bson:"endDateTime,omitempty"`

	// Start of the time period, using IETC-RFC-3339 format. If you define a start, you must also define an end
	StartDateTime *time.Time `json:"startDateTime,omitempty" bson:"startDateTime,omitempty"`
}

/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package main

import (
	"resource-inventory/src/routes"
)

func main() {
	var s routes.Routes
	s.StartGin()
}

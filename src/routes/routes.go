/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"resource-inventory/src/handlers"
)

type Routes struct {
}

func (c Routes) StartGin() {
	r := gin.Default()
	resourceApi := r.Group("/tmf-api/resourceInventoryManagement/v4/")
	{
		resourceApi.GET("/resource", handlers.ListResource)
		resourceApi.POST("/resource", handlers.CreateResource)
		resourceApi.GET("/resource/:id", handlers.RetrieveResource)
		resourceApi.PATCH("/resource/:id", handlers.PatchResource)
		resourceApi.DELETE("/resource/:id", handlers.DeleteResource)
	}
	err := r.Run()
	if err != nil {
		return
	}
}

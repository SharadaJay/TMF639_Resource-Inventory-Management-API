/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"resource-inventory/src/handlers"
)

type Routes struct {
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (c Routes) StartGin() {
	r := gin.Default()
	r.Use(CORSMiddleware())
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
		log.Fatalf("Failed to run gin router")
	}
}

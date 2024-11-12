/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-inventory/src/models"
	"resource-inventory/src/services"
)

const ResourceCollection = "Resources"

func ListResource(c *gin.Context) {
	fields := c.Query("fields")
	limit := c.Query("limit")
	offset := c.Query("offset")
	filterByValString := c.Query("filterByValue")

	resources, err := services.ListResources(fields, limit, offset, filterByValString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching resources",
		})
		return
	}

	c.JSON(http.StatusOK, resources)
}

func CreateResource(c *gin.Context) {
	var resource models.Resource
	if err := c.BindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	createdResource, err := services.CreateResource(resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating resource"})
		return
	}
	c.JSON(http.StatusCreated, createdResource)
}

func RetrieveResource(c *gin.Context) {
	id := c.Param("id")
	resource, err := services.RetrieveResource(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		return
	}
	c.JSON(http.StatusOK, resource)
}

func PatchResource(c *gin.Context) {
	id := c.Param("id")
	var update models.Resource
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	updatedResource, err := services.PatchResource(id, update)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		return
	}

	c.JSON(http.StatusOK, updatedResource)
}

func DeleteResource(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteResource(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

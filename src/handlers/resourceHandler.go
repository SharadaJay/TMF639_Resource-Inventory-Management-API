/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"resource-inventory/src/models"
	"resource-inventory/src/services"
)

func ListResource(c *gin.Context) {
	log.Debug("ListResource handler started")

	fields := c.Query("fields")
	limit := c.Query("limit")
	offset := c.Query("offset")
	filterByValString := c.Query("filterByValue")

	log.Debugf("Received query parameters - fields: %s, limit: %s, offset: %s, filterByValue: %s", fields, limit, offset, filterByValString)

	resources, err := services.ListResources(fields, limit, offset, filterByValString)
	if err != nil {
		log.Error("Error fetching resources: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching resources",
		})
		return
	}

	log.Debug("ListResource handler completed")
	c.JSON(http.StatusOK, resources)
}

func CreateResource(c *gin.Context) {
	log.Debug("CreateResource handler started")

	var resource models.Resource
	if err := c.BindJSON(&resource); err != nil {
		log.Error("Failed to bind JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	createdResource, err := services.CreateResource(resource)
	if err != nil {
		log.Error("Failed to create resource: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating resource"})
		return
	}

	log.Debug("CreateResource handler completed")
	c.JSON(http.StatusCreated, createdResource)
}

func RetrieveResource(c *gin.Context) {
	log.Debug("RetrieveResource handler started")

	id := c.Param("id")
	resource, err := services.RetrieveResource(id)
	if err != nil {
		log.Error("Failed to retrieve resource: ", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		return
	}

	log.Debug("RetrieveResource handler completed")
	c.JSON(http.StatusOK, resource)
}

func PatchResource(c *gin.Context) {
	log.Debug("PatchResource handler started")

	id := c.Param("id")
	var update models.Resource
	if err := c.BindJSON(&update); err != nil {
		log.Error("Failed to bind JSON for update: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	updatedResource, err := services.PatchResource(id, update)
	if err != nil {
		log.Error("Failed to update resource: ", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		return
	}

	log.Debug("PatchResource handler completed")
	c.JSON(http.StatusOK, updatedResource)
}

func DeleteResource(c *gin.Context) {
	log.Debug("DeleteResource handler started")

	id := c.Param("id")
	err := services.DeleteResource(id)
	if err != nil {
		log.Error("Failed to delete resource: ", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		return
	}

	log.Debug("DeleteResource handler completed")
	c.JSON(http.StatusNoContent, gin.H{})
}

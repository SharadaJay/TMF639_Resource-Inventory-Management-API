/*
 * Resource Inventory Management
 * API version: 4.0.0
 */

package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListResource(c *gin.Context) {
	c.JSON(http.StatusOK, "Resources Listed")
}

func RetrieveResource(c *gin.Context) {
	c.JSON(http.StatusOK, "Resource Retrieved")
}

func CreateResource(c *gin.Context) {
	c.JSON(http.StatusCreated, "Resource Created")
}

func PatchResource(c *gin.Context) {
	c.JSON(http.StatusOK, "Resource Patched")
}

func DeleteResource(c *gin.Context) {
	c.JSON(http.StatusNoContent, "Resources Deleted")
}

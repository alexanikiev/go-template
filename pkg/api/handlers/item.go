package handlers

import (
	"net/http"
	"strconv"

	"gotemplate/pkg/api/models"

	"github.com/gin-gonic/gin"
)

// ReadRoot handles the root endpoint
func ReadRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
}

// ReadItem retrieves an item by its ID
func ReadItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item_id": itemID})
}

// CreateItem creates a new item
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": item.Name, "description": item.Description})
}

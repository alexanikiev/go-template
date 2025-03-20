package routes

import (
	"gotemplate/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", handlers.ReadRoot)
	router.GET("/items/:item_id", handlers.ReadItem)
	router.POST("/items", handlers.CreateItem)

	return router
}

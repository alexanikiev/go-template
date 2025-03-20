// pkg/solver/api.go
package api

import (
	"gotemplate/pkg/api/routes"
	"log"
)

// StartServer runs the API server
func StartServer() {
	router := routes.SetupRouter()

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

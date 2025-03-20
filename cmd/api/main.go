// cmd/api/main.go
package main

import (
	"gotemplate/pkg/api"
	"log"
)

func main() {
	log.Println("Starting API server...")
	api.StartServer()
}

// cmd/worker/main.go
package main

import (
	"gotemplate/pkg/worker"
	"log"
)

func main() {
	log.Println("Starting Worker...")
	worker.RunWorker()
	log.Println("Worker finished execution.")
}

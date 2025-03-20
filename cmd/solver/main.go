// cmd/solver/main.go
package main

import (
	"fmt"
	"gotemplate/pkg/solver"
	"log"
)

func main() {
	// Create a Solver instance
	s := solver.NewSolver("a", "b")

	x, y := 3, 4
	log.Printf("Calling Solve method with x=%d, y=%d", x, y)

	// Run solve function
	result := s.Solve(x, y)

	// Print result
	log.Printf("Computation successful. Result: %d", result)
	fmt.Printf("Result: %d\n", result)
}

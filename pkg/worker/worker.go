// pkg/solver/worker.go
package worker

import (
	"fmt"
	"gotemplate/pkg/solver"
	"sync"
)

// asyncSolve runs the Solve method asynchronously using a Goroutine
func asyncSolve(solverInstance *solver.Solver, x, y int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when function exits

	result := solverInstance.Solve(x, y)
	fmt.Printf("Async result: %d\n", result)
}

// RunWorker initializes a solver instance and runs asyncSolve in a Goroutine
func RunWorker() {
	solverInstance := solver.NewSolver("a", "b")

	var wg sync.WaitGroup
	wg.Add(1) // Increment counter

	go asyncSolve(solverInstance, 1, 2, &wg)

	wg.Wait() // Wait for Goroutines to finish
}

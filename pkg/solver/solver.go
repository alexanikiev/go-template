// pkg/solver/solver.go
package solver

import (
	"errors"
)

// Solver represents a solver for a sample problem
type Solver struct {
	A string
	B string
}

// NewSolver is a constructor function to create a Solver instance
func NewSolver(a, b string) *Solver {
	return &Solver{A: a, B: b}
}

// calculate is a private method (lowercase, not exported)
func (s *Solver) calculate(x, y int) int {
	return x * y
}

// Solve is a public method (uppercase, exported)
func (s *Solver) Solve(x, y int) int {
	return s.calculate(x, y)
}

// Error is a public method that raises an error
func (s *Solver) Error() error {
	return errors.New("this is an error message")
}

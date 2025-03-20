package solver

import (
	"testing"
)

// setUp initializes a Solver instance for each test case
func setUp() *Solver {
	return NewSolver("a", "b")
}

// TestSolve checks if Solve works for normal values
func TestSolve(t *testing.T) {
	solver := setUp()
	result := solver.Solve(3, 4)
	expected := 12

	if result != expected {
		t.Errorf("Solve(3, 4) = %d; want %d", result, expected)
	}
}

// TestSolveWithZero checks if Solve works when one number is zero
func TestSolveWithZero(t *testing.T) {
	solver := setUp()
	result := solver.Solve(0, 5)
	expected := 0

	if result != expected {
		t.Errorf("Solve(0, 5) = %d; want %d", result, expected)
	}
}

// TestSolveWithNegativeNumbers checks if Solve handles negative numbers correctly
func TestSolveWithNegativeNumbers(t *testing.T) {
	solver := setUp()

	tests := []struct {
		x, y     int
		expected int
	}{
		{-3, 4, -12},
		{-3, -4, 12},
	}

	for _, test := range tests {
		result := solver.Solve(test.x, test.y)
		if result != test.expected {
			t.Errorf("Solve(%d, %d) = %d; want %d", test.x, test.y, result, test.expected)
		}
	}
}

// TestSolveWithLargeNumbers checks if Solve handles large numbers correctly
func TestSolveWithLargeNumbers(t *testing.T) {
	solver := setUp()
	result := solver.Solve(1000000, 1000000)
	expected := 1000000000000

	if result != expected {
		t.Errorf("Solve(1000000, 1000000) = %d; want %d", result, expected)
	}
}

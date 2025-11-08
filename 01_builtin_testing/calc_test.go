package calc_test

import (
	"testing"

	"github.com/lirany1/go-testing-framework-examples/01_builtin_testing/calc"
)

// TestSum demonstrates a basic test using Go's built-in testing package.
func TestSum(t *testing.T) {
	result := calc.Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Sum(2, 3) = %d; want %d", result, expected)
	}
}

// TestSum_TableDriven demonstrates table-driven tests - a common Go testing pattern.
// This approach allows testing multiple scenarios with minimal code duplication.
func TestSum_TableDriven(t *testing.T) {
	// Define test cases as a slice of structs
	tests := []struct {
		name     string // descriptive name for the test case
		a, b     int    // input parameters
		expected int    // expected result
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -2, -3},
		{"with zero", 0, 5, 5},
		{"both zero", 0, 0, 0},
		{"large numbers", 1000, 2000, 3000},
	}

	// Iterate through test cases
	for _, tt := range tests {
		// t.Run creates a subtest for each case
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Sum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMultiply demonstrates testing another function.
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 6},
		{"by zero", 5, 0, 0},
		{"negative numbers", -2, 3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestDivide demonstrates testing with edge cases (division by zero).
func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"normal division", 6, 2, 3},
		{"divide by zero", 5, 0, 0}, // our implementation returns 0 for division by zero
		{"negative result", -6, 2, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Divide(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// BenchmarkSum demonstrates Go's built-in benchmarking capabilities.
// Run with: go test -bench=.
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc.Sum(2, 3)
	}
}

// TestSum_Parallel demonstrates parallel test execution.
// Parallel tests can significantly speed up test suites.
func TestSum_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"case 1", 1, 1, 2},
		{"case 2", 2, 2, 4},
		{"case 3", 3, 3, 6},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel tests
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // marks this test to run in parallel
			result := calc.Sum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

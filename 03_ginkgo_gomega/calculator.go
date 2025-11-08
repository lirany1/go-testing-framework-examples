package ginkgo_gomega

// Calculator provides basic arithmetic operations.
type Calculator struct{}

// NewCalculator creates a new Calculator instance.
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add returns the sum of two integers.
func (c *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers.
func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers.
func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers.
// Returns an error if attempting to divide by zero.
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// ErrDivisionByZero is returned when dividing by zero.
var ErrDivisionByZero = Error("cannot divide by zero")

// Error is a simple error type.
type Error string

func (e Error) Error() string {
	return string(e)
}

// IsPositive checks if a number is positive.
func (c *Calculator) IsPositive(n int) bool {
	return n > 0
}

// Sum returns the sum of a slice of integers.
func (c *Calculator) Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

package calc

// Sum returns the sum of two integers.
// This is a simple function used to demonstrate Go's built-in testing package.
func Sum(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the division of two integers.
// Returns 0 if attempting to divide by zero.
func Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

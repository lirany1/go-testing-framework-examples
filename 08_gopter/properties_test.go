package gopter

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

// Sum returns the sum of two integers.
func Sum(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// Abs returns the absolute value of an integer.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Reverse reverses a slice.
func Reverse(s []int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

// TestSumProperties demonstrates property-based testing for addition.
func TestSumProperties(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100 // Run 100 random tests

	properties := gopter.NewProperties(parameters)

	// Property 1: Addition is commutative (a + b == b + a)
	properties.Property("addition is commutative", prop.ForAll(
		func(a, b int) bool {
			return Sum(a, b) == Sum(b, a)
		},
		gen.Int(), // Generate random int for a
		gen.Int(), // Generate random int for b
	))

	// Property 2: Zero is the identity element (a + 0 == a)
	properties.Property("zero is the identity element", prop.ForAll(
		func(a int) bool {
			return Sum(a, 0) == a
		},
		gen.Int(),
	))

	// Property 3: Addition is associative ((a + b) + c == a + (b + c))
	properties.Property("addition is associative", prop.ForAll(
		func(a, b, c int) bool {
			return Sum(Sum(a, b), c) == Sum(a, Sum(b, c))
		},
		gen.Int(), gen.Int(), gen.Int(),
	))

	// Property 4: Adding positive numbers increases the value
	properties.Property("adding positive numbers increases value", prop.ForAll(
		func(a, b int) bool {
			// Only test when b is positive
			if b > 0 {
				return Sum(a, b) > a
			}
			return true
		},
		gen.Int(),
		gen.IntRange(1, 1000), // Generate only positive numbers
	))

	properties.TestingRun(t)
}

// TestMultiplyProperties demonstrates more complex properties.
func TestMultiplyProperties(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	// Multiplication is commutative
	properties.Property("multiplication is commutative", prop.ForAll(
		func(a, b int) bool {
			return Multiply(a, b) == Multiply(b, a)
		},
		gen.Int(), gen.Int(),
	))

	// One is the identity element
	properties.Property("one is the identity element", prop.ForAll(
		func(a int) bool {
			return Multiply(a, 1) == a
		},
		gen.Int(),
	))

	// Zero property
	properties.Property("anything times zero is zero", prop.ForAll(
		func(a int) bool {
			return Multiply(a, 0) == 0
		},
		gen.Int(),
	))

	// Multiplication is associative
	properties.Property("multiplication is associative", prop.ForAll(
		func(a, b, c int) bool {
			return Multiply(Multiply(a, b), c) == Multiply(a, Multiply(b, c))
		},
		gen.IntRange(-100, 100), // Limit range to avoid overflow
		gen.IntRange(-100, 100),
		gen.IntRange(-100, 100),
	))

	// Distributive property
	properties.Property("multiplication distributes over addition", prop.ForAll(
		func(a, b, c int) bool {
			return Multiply(a, Sum(b, c)) == Sum(Multiply(a, b), Multiply(a, c))
		},
		gen.IntRange(-100, 100),
		gen.IntRange(-100, 100),
		gen.IntRange(-100, 100),
	))

	properties.TestingRun(t)
}

// TestAbsProperties demonstrates properties of absolute value.
func TestAbsProperties(t *testing.T) {
	properties := gopter.NewProperties(gopter.DefaultTestParameters())

	// Abs is always non-negative
	properties.Property("abs is always non-negative", prop.ForAll(
		func(x int) bool {
			return Abs(x) >= 0
		},
		gen.Int(),
	))

	// Abs is idempotent: abs(abs(x)) == abs(x)
	properties.Property("abs is idempotent", prop.ForAll(
		func(x int) bool {
			return Abs(Abs(x)) == Abs(x)
		},
		gen.Int(),
	))

	// For positive numbers, abs(x) == x
	properties.Property("abs of positive number is itself", prop.ForAll(
		func(x int) bool {
			if x >= 0 {
				return Abs(x) == x
			}
			return true
		},
		gen.Int(),
	))

	// For negative numbers, abs(x) == -x
	properties.Property("abs of negative number is its negation", prop.ForAll(
		func(x int) bool {
			if x < 0 {
				return Abs(x) == -x
			}
			return true
		},
		gen.Int(),
	))

	properties.TestingRun(t)
}

// TestReverseProperties demonstrates properties of slice reversal.
func TestReverseProperties(t *testing.T) {
	properties := gopter.NewProperties(gopter.DefaultTestParameters())

	// Reversing twice returns original
	properties.Property("reversing twice returns original", prop.ForAll(
		func(s []int) bool {
			reversed := Reverse(s)
			doubleReversed := Reverse(reversed)

			if len(s) != len(doubleReversed) {
				return false
			}

			for i := range s {
				if s[i] != doubleReversed[i] {
					return false
				}
			}
			return true
		},
		gen.SliceOf(gen.Int()),
	))

	// Length is preserved
	properties.Property("length is preserved", prop.ForAll(
		func(s []int) bool {
			return len(s) == len(Reverse(s))
		},
		gen.SliceOf(gen.Int()),
	))

	// First becomes last
	properties.Property("first element becomes last", prop.ForAll(
		func(s []int) bool {
			if len(s) == 0 {
				return true
			}
			reversed := Reverse(s)
			return s[0] == reversed[len(reversed)-1]
		},
		gen.SliceOf(gen.Int()),
	))

	// Last becomes first
	properties.Property("last element becomes first", prop.ForAll(
		func(s []int) bool {
			if len(s) == 0 {
				return true
			}
			reversed := Reverse(s)
			return s[len(s)-1] == reversed[0]
		},
		gen.SliceOf(gen.Int()),
	))

	properties.TestingRun(t)
}

// TestStringProperties demonstrates property testing with strings.
func TestStringProperties(t *testing.T) {
	properties := gopter.NewProperties(gopter.DefaultTestParameters())

	// Length of concatenation
	properties.Property("length of concat equals sum of lengths", prop.ForAll(
		func(s1, s2 string) bool {
			concat := s1 + s2
			return len(concat) == len(s1)+len(s2)
		},
		gen.AnyString(),
		gen.AnyString(),
	))

	properties.TestingRun(t)
}

// TestCustomGenerators demonstrates creating custom generators.
func TestCustomGenerators(t *testing.T) {
	properties := gopter.NewProperties(gopter.DefaultTestParameters())

	// Custom generator for even numbers
	evenGen := gen.Int().SuchThat(func(x int) bool {
		return x%2 == 0
	})

	properties.Property("generated numbers are even", prop.ForAll(
		func(x int) bool {
			return x%2 == 0
		},
		evenGen,
	))

	// Custom generator for non-empty strings
	nonEmptyStringGen := gen.AnyString().SuchThat(func(s string) bool {
		return len(s) > 0
	})

	properties.Property("generated strings are non-empty", prop.ForAll(
		func(s string) bool {
			return len(s) > 0
		},
		nonEmptyStringGen,
	))

	properties.TestingRun(t)
}

package goconvey

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Sum returns the sum of two integers.
func Sum(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// Divide divides two integers.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// ErrDivisionByZero is returned when dividing by zero.
type errDivisionByZero struct{}

func (e errDivisionByZero) Error() string {
	return "division by zero"
}

var ErrDivisionByZero = errDivisionByZero{}

// TestSum demonstrates GoConvey's BDD-style syntax.
// Run with: go test
// Or better: goconvey (opens Web UI at http://localhost:8080)
func TestSum(t *testing.T) {
	// Convey starts a top-level test scope
	Convey("Sum function", t, func() {

		Convey("should add positive numbers correctly", func() {
			result := Sum(2, 3)
			So(result, ShouldEqual, 5)
		})

		Convey("should handle negative numbers", func() {
			result := Sum(-2, -3)
			So(result, ShouldEqual, -5)
		})

		Convey("should handle zero", func() {
			result := Sum(0, 5)
			So(result, ShouldEqual, 5)
		})

		// Nested Convey blocks for better organization
		Convey("with nested contexts", func() {
			Convey("when adding large numbers", func() {
				result := Sum(1000, 2000)
				So(result, ShouldEqual, 3000)
			})

			Convey("when result is negative", func() {
				result := Sum(-10, 5)
				So(result, ShouldEqual, -5)
			})
		})
	})
}

// TestMultiply demonstrates multiple assertions in nested contexts.
func TestMultiply(t *testing.T) {
	Convey("Multiply function", t, func() {

		Convey("Basic multiplication", func() {
			Convey("should multiply positive numbers", func() {
				So(Multiply(3, 4), ShouldEqual, 12)
			})

			Convey("should handle multiplication by zero", func() {
				So(Multiply(5, 0), ShouldEqual, 0)
				So(Multiply(0, 5), ShouldEqual, 0)
			})

			Convey("should handle negative numbers", func() {
				So(Multiply(-3, 4), ShouldEqual, -12)
				So(Multiply(-3, -4), ShouldEqual, 12)
			})
		})

		Convey("Edge cases", func() {
			Convey("should handle one as multiplier", func() {
				So(Multiply(5, 1), ShouldEqual, 5)
			})

			Convey("should be commutative", func() {
				result1 := Multiply(3, 7)
				result2 := Multiply(7, 3)
				So(result1, ShouldEqual, result2)
			})
		})
	})
}

// TestDivide demonstrates error handling with GoConvey.
func TestDivide(t *testing.T) {
	Convey("Divide function", t, func() {

		Convey("Valid division", func() {
			Convey("should divide evenly", func() {
				result, err := Divide(6, 2)
				So(err, ShouldBeNil)
				So(result, ShouldEqual, 3)
			})

			Convey("should handle negative numbers", func() {
				result, err := Divide(-6, 2)
				So(err, ShouldBeNil)
				So(result, ShouldEqual, -3)
			})

			Convey("should handle integer division", func() {
				result, err := Divide(7, 2)
				So(err, ShouldBeNil)
				So(result, ShouldEqual, 3) // integer division truncates
			})
		})

		Convey("Division by zero", func() {
			Convey("should return an error", func() {
				result, err := Divide(5, 0)
				So(err, ShouldNotBeNil)
				So(err, ShouldResemble, ErrDivisionByZero)
				So(result, ShouldEqual, 0)
			})

			Convey("error message should be descriptive", func() {
				_, err := Divide(10, 0)
				So(err.Error(), ShouldContainSubstring, "division by zero")
			})
		})
	})
}

// TestAdvancedMatchers demonstrates various GoConvey assertions.
func TestAdvancedMatchers(t *testing.T) {
	Convey("Advanced GoConvey matchers", t, func() {

		Convey("Numeric comparisons", func() {
			value := 10
			So(value, ShouldBeGreaterThan, 5)
			So(value, ShouldBeLessThan, 15)
			So(value, ShouldBeGreaterThanOrEqualTo, 10)
			So(value, ShouldBeLessThanOrEqualTo, 10)
		})

		Convey("String operations", func() {
			str := "Hello, World!"
			So(str, ShouldContainSubstring, "World")
			So(str, ShouldStartWith, "Hello")
			So(str, ShouldEndWith, "!")
			So(str, ShouldNotBeEmpty)
		})

		Convey("Collection assertions", func() {
			slice := []int{1, 2, 3, 4, 5}
			So(slice, ShouldHaveLength, 5)
			So(slice, ShouldContain, 3)
			So(slice, ShouldNotContain, 10)
		})

		Convey("Boolean checks", func() {
			So(true, ShouldBeTrue)
			So(false, ShouldBeFalse)
			So(Sum(2, 2) == 4, ShouldBeTrue)
		})

		Convey("Nil checks", func() {
			var ptr *int
			So(ptr, ShouldBeNil)

			value := 42
			ptr = &value
			So(ptr, ShouldNotBeNil)
		})

		Convey("Type checks", func() {
			var x interface{} = 42
			So(x, ShouldHaveSameTypeAs, 0)
		})
	})
}

// TestWithReset demonstrates Reset for cleanup between tests.
func TestWithReset(t *testing.T) {
	Convey("Tests with cleanup", t, func() {
		counter := 0

		// Reset runs after each Convey at this level
		Reset(func() {
			counter = 0
		})

		Convey("First test modifies counter", func() {
			counter++
			So(counter, ShouldEqual, 1)
		})

		Convey("Second test starts with clean counter", func() {
			// Thanks to Reset, counter is back to 0
			So(counter, ShouldEqual, 0)
			counter += 5
			So(counter, ShouldEqual, 5)
		})
	})
}

// TestComplexScenario demonstrates a more realistic test scenario.
func TestComplexScenario(t *testing.T) {
	Convey("Given a calculator", t, func() {

		Convey("When performing multiple operations", func() {
			// (10 + 5) * 2 = 30
			result := Sum(10, 5)
			result = Multiply(result, 2)

			Convey("The result should be correct", func() {
				So(result, ShouldEqual, 30)
			})

			Convey("The result should be positive", func() {
				So(result, ShouldBeGreaterThan, 0)
			})

			Convey("When dividing the result", func() {
				quotient, err := Divide(result, 3)

				So(err, ShouldBeNil)
				So(quotient, ShouldEqual, 10)
			})
		})
	})
}

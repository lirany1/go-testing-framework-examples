package ginkgo_gomega_test

import (
	ginkgo_gomega "github.com/lirany1/go-testing-framework-examples/03_ginkgo_gomega"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Demonstrate BDD-style testing with Ginkgo and Gomega.
// Describe/Context/It provide a natural way to organize tests.
var _ = Describe("Calculator", func() {
	var calc *ginkgo_gomega.Calculator

	// BeforeEach runs before each test (It block)
	BeforeEach(func() {
		calc = ginkgo_gomega.NewCalculator()
	})

	// Describe groups related tests together
	Describe("Addition", func() {
		// Context describes different scenarios
		Context("when adding positive numbers", func() {
			It("should return the correct sum", func() {
				result := calc.Add(2, 3)
				Expect(result).To(Equal(5))
			})

			It("should be commutative", func() {
				result1 := calc.Add(2, 3)
				result2 := calc.Add(3, 2)
				Expect(result1).To(Equal(result2))
			})
		})

		Context("when adding negative numbers", func() {
			It("should handle negative operands", func() {
				result := calc.Add(-2, -3)
				Expect(result).To(Equal(-5))
			})

			It("should handle mixed signs", func() {
				result := calc.Add(-5, 3)
				Expect(result).To(Equal(-2))
			})
		})

		Context("when adding with zero", func() {
			It("should return the other number", func() {
				Expect(calc.Add(0, 5)).To(Equal(5))
				Expect(calc.Add(5, 0)).To(Equal(5))
				Expect(calc.Add(0, 0)).To(Equal(0))
			})
		})
	})

	Describe("Subtraction", func() {
		Context("when subtracting positive numbers", func() {
			It("should return the correct difference", func() {
				result := calc.Subtract(5, 3)
				Expect(result).To(Equal(2))
			})

			It("should handle negative results", func() {
				result := calc.Subtract(3, 5)
				Expect(result).To(Equal(-2))
			})
		})

		Context("when subtracting negative numbers", func() {
			It("should handle double negatives", func() {
				result := calc.Subtract(5, -3)
				Expect(result).To(Equal(8))
			})
		})
	})

	Describe("Multiplication", func() {
		It("should multiply positive numbers correctly", func() {
			result := calc.Multiply(3, 4)
			Expect(result).To(Equal(12))
		})

		It("should handle multiplication by zero", func() {
			Expect(calc.Multiply(5, 0)).To(Equal(0))
			Expect(calc.Multiply(0, 5)).To(Equal(0))
		})

		It("should handle negative numbers", func() {
			Expect(calc.Multiply(-3, 4)).To(Equal(-12))
			Expect(calc.Multiply(-3, -4)).To(Equal(12))
		})
	})

	Describe("Division", func() {
		Context("when dividing valid numbers", func() {
			It("should return the correct quotient", func() {
				result, err := calc.Divide(6, 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(3))
			})

			It("should handle negative numbers", func() {
				result, err := calc.Divide(-6, 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(-3))
			})
		})

		Context("when dividing by zero", func() {
			It("should return an error", func() {
				result, err := calc.Divide(5, 0)

				// Multiple expectations in one It block
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(ginkgo_gomega.ErrDivisionByZero))
				Expect(result).To(Equal(0))
			})

			It("should have a descriptive error message", func() {
				_, err := calc.Divide(10, 0)
				Expect(err.Error()).To(ContainSubstring("divide by zero"))
			})
		})
	})

	Describe("IsPositive", func() {
		// Table-driven tests can also be done with Ginkgo
		DescribeTable("checking if number is positive",
			func(num int, expected bool) {
				result := calc.IsPositive(num)
				Expect(result).To(Equal(expected))
			},
			Entry("positive number", 5, true),
			Entry("negative number", -5, false),
			Entry("zero", 0, false),
			Entry("large positive", 1000, true),
			Entry("large negative", -1000, false),
		)
	})

	Describe("Sum with slice", func() {
		Context("when summing a slice of numbers", func() {
			It("should return correct sum for positive numbers", func() {
				numbers := []int{1, 2, 3, 4, 5}
				result := calc.Sum(numbers)
				Expect(result).To(Equal(15))
			})

			It("should handle empty slice", func() {
				result := calc.Sum([]int{})
				Expect(result).To(Equal(0))
			})

			It("should handle slice with negative numbers", func() {
				numbers := []int{-1, -2, -3}
				result := calc.Sum(numbers)
				Expect(result).To(Equal(-6))
			})

			It("should handle mixed positive and negative", func() {
				numbers := []int{10, -5, 3, -2}
				result := calc.Sum(numbers)
				Expect(result).To(Equal(6))
			})
		})

		Context("demonstrating Gomega matchers", func() {
			It("should work with numeric comparisons", func() {
				result := calc.Sum([]int{1, 2, 3})

				Expect(result).To(BeNumerically(">", 5))
				Expect(result).To(BeNumerically("<=", 6))
				Expect(result).To(BeNumerically("==", 6))
			})

			It("should demonstrate collection matchers", func() {
				numbers := []int{1, 2, 3, 4, 5}

				Expect(numbers).To(HaveLen(5))
				Expect(numbers).To(ContainElement(3))
				Expect(numbers).To(ConsistOf(1, 2, 3, 4, 5))
				Expect(numbers).NotTo(ContainElement(10))
			})
		})
	})

	// Demonstrating focused specs (useful during development)
	// Uncomment FDescribe or FIt to run only those tests
	/*
		FDescribe("Focused test", func() {
			FIt("runs only this test", func() {
				Expect(true).To(BeTrue())
			})
		})
	*/

	// Demonstrating pending tests
	PDescribe("Future feature", func() {
		It("will be implemented later", func() {
			// This test is marked as pending and won't run
		})
	})
})

// Demonstrate nested contexts for complex scenarios
var _ = Describe("Calculator Advanced Scenarios", func() {
	var calc *ginkgo_gomega.Calculator

	BeforeEach(func() {
		calc = ginkgo_gomega.NewCalculator()
	})

	Describe("Complex calculations", func() {
		Context("when performing multiple operations", func() {
			It("should maintain accuracy", func() {
				// (10 + 5) * 2 - 6 = 24
				result := calc.Add(10, 5)
				result = calc.Multiply(result, 2)
				result = calc.Subtract(result, 6)

				Expect(result).To(Equal(24))
			})
		})

		Context("when chaining operations with error handling", func() {
			It("should handle errors gracefully", func() {
				result := calc.Multiply(10, 2)          // 20
				quotient, err := calc.Divide(result, 4) // 5

				Expect(err).NotTo(HaveOccurred())
				Expect(quotient).To(Equal(5))
			})
		})
	})
})

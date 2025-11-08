//go:build gauge || ignore
// +build gauge ignore

package main

import (
	"fmt"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/getgauge-contrib/gauge-go/models"
	"github.com/lirany1/go-testing-framework-examples/07_gauge/testsuit"
)

var _ = gauge.Step("Initialize calculator", func() {
	testsuit.Calc = &testsuit.Calculator{}
	testsuit.Numbers = []int{}
	testsuit.Result = 0
	testsuit.Error = ""
})

var _ = gauge.Step("Enter number <number>", func(number int) {
	testsuit.Numbers = append(testsuit.Numbers, number)
})

var _ = gauge.Step("Press <operation> button", func(operation string) {
	if len(testsuit.Numbers) < 2 {
		testsuit.Error = "need at least 2 numbers"
		return
	}

	a, b := testsuit.Numbers[0], testsuit.Numbers[1]

	switch operation {
	case "add":
		testsuit.Result = testsuit.Calc.Add(a, b)
	case "subtract":
		testsuit.Result = testsuit.Calc.Subtract(a, b)
	case "multiply":
		testsuit.Result = testsuit.Calc.Multiply(a, b)
	case "divide":
		result, err := testsuit.Calc.Divide(a, b)
		if err != nil {
			testsuit.Error = err.Error()
		} else {
			testsuit.Result = result
		}
	default:
		testsuit.Error = "unknown operation: " + operation
	}
})

var _ = gauge.Step("Result should be <expected>", func(expected int) {
	if testsuit.Result != expected {
		gauge.WriteMessage("Expected %d but got %d", expected, testsuit.Result)
		panic(fmt.Sprintf("Expected %d but got %d", expected, testsuit.Result))
	}
})

var _ = gauge.Step("Should see error <message>", func(message string) {
	if testsuit.Error != message {
		panic(fmt.Sprintf("Expected error '%s' but got '%s'", message, testsuit.Error))
	}
})

var _ = gauge.Step("Perform calculation with data table", func(table *models.Table) {
	for _, row := range table.Rows {
		operation := row.Cells[0]
		first := parseInt(row.Cells[1])
		second := parseInt(row.Cells[2])
		expectedResult := parseInt(row.Cells[3])

		testsuit.Numbers = []int{first, second}

		switch operation {
		case "add":
			testsuit.Result = testsuit.Calc.Add(first, second)
		case "subtract":
			testsuit.Result = testsuit.Calc.Subtract(first, second)
		case "multiply":
			testsuit.Result = testsuit.Calc.Multiply(first, second)
		case "divide":
			testsuit.Result, _ = testsuit.Calc.Divide(first, second)
		}

		if testsuit.Result != expectedResult {
			panic(fmt.Sprintf("For %s(%d, %d): expected %d but got %d",
				operation, first, second, expectedResult, testsuit.Result))
		}
	}
})

func parseInt(s string) int {
	var i int
	_, _ = fmt.Sscanf(s, "%d", &i) // Error ignored - returns 0 on failure which is acceptable for tests
	return i
}

// Calculator provides basic arithmetic operations.
type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

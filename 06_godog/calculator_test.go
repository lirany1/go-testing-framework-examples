package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
)

// calculatorContext holds the state for calculator scenarios.
type calculatorContext struct {
	calculator Calculator
	numbers    []int
	result     int
	err        error
}

// Calculator provides arithmetic operations.
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

// Step definitions

func (cc *calculatorContext) aCalculator() error {
	cc.calculator = Calculator{}
	cc.numbers = []int{}
	cc.result = 0
	cc.err = nil
	return nil
}

func (cc *calculatorContext) iHaveEnteredIntoTheCalculator(number int) error {
	cc.numbers = append(cc.numbers, number)
	return nil
}

func (cc *calculatorContext) iPressAdd() error {
	if len(cc.numbers) < 2 {
		return fmt.Errorf("need at least 2 numbers")
	}
	cc.result = cc.calculator.Add(cc.numbers[0], cc.numbers[1])
	return nil
}

func (cc *calculatorContext) iPressSubtract() error {
	if len(cc.numbers) < 2 {
		return fmt.Errorf("need at least 2 numbers")
	}
	cc.result = cc.calculator.Subtract(cc.numbers[0], cc.numbers[1])
	return nil
}

func (cc *calculatorContext) iPressMultiply() error {
	if len(cc.numbers) < 2 {
		return fmt.Errorf("need at least 2 numbers")
	}
	cc.result = cc.calculator.Multiply(cc.numbers[0], cc.numbers[1])
	return nil
}

func (cc *calculatorContext) iPressDivide() error {
	if len(cc.numbers) < 2 {
		return fmt.Errorf("need at least 2 numbers")
	}
	cc.result, cc.err = cc.calculator.Divide(cc.numbers[0], cc.numbers[1])
	return nil
}

func (cc *calculatorContext) iPressOperation(operation string) error {
	switch operation {
	case "add":
		return cc.iPressAdd()
	case "subtract":
		return cc.iPressSubtract()
	case "multiply":
		return cc.iPressMultiply()
	case "divide":
		return cc.iPressDivide()
	default:
		return fmt.Errorf("unknown operation: %s", operation)
	}
}

func (cc *calculatorContext) theResultShouldBeOnTheScreen(expected int) error {
	if cc.result != expected {
		return fmt.Errorf("expected %d, got %d", expected, cc.result)
	}
	return nil
}

func (cc *calculatorContext) iShouldSeeAnErrorMessage(expectedMsg string) error {
	if cc.err == nil {
		return fmt.Errorf("expected error with message '%s', but got no error", expectedMsg)
	}
	if cc.err.Error() != expectedMsg {
		return fmt.Errorf("expected error '%s', got '%s'", expectedMsg, cc.err.Error())
	}
	return nil
}

// InitializeScenario registers step definitions.
func InitializeScenario(sc *godog.ScenarioContext) {
	cc := &calculatorContext{}

	// Register step definitions
	sc.Step(`^a calculator$`, cc.aCalculator)
	sc.Step(`^I have entered (\d+) into the calculator$`, cc.iHaveEnteredIntoTheCalculator)
	sc.Step(`^I press add$`, cc.iPressAdd)
	sc.Step(`^I press subtract$`, cc.iPressSubtract)
	sc.Step(`^I press multiply$`, cc.iPressMultiply)
	sc.Step(`^I press divide$`, cc.iPressDivide)
	sc.Step(`^I press (.+)$`, cc.iPressOperation)
	sc.Step(`^the result should be (\d+) on the screen$`, cc.theResultShouldBeOnTheScreen)
	sc.Step(`^I should see an error message "([^"]*)"$`, cc.iShouldSeeAnErrorMessage)
}

// TestFeatures runs the Godog test suite.
func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func TestMain(m *testing.M) {
	status := m.Run()
	os.Exit(status)
}

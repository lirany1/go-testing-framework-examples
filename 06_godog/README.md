# Godog

Godog is the official Cucumber implementation for Go, enabling BDD with Gherkin syntax. Write test scenarios in plain English (Given/When/Then) and map them to Go step definitions.

## 📦 Installation

```bash
go get github.com/cucumber/godog
```

## 🎯 Features

- **Gherkin Syntax**: Write scenarios in natural language
- **Cucumber Compatible**: Standard BDD tool
- **Acceptance Testing**: Perfect for end-to-end tests
- **Multiple Formatters**: Pretty, JSON, JUnit output
- **Tags**: Run specific scenarios
- **Hooks**: BeforeScenario, AfterScenario, etc.

## 📖 Usage

### Feature File (.feature)

```gherkin
Feature: Addition
  Scenario: Add two positive numbers
    Given I have numbers 2 and 3
    When I add them
    Then the result should be 5
```

### Step Definitions (Go)

```go
func iHaveNumbers(a, b int) error {
    ctx.a = a
    ctx.b = b
    return nil
}

func iAddThem() error {
    ctx.result = ctx.a + ctx.b
    return nil
}

func theResultShouldBe(expected int) error {
    if ctx.result != expected {
        return fmt.Errorf("expected %d, got %d", expected, ctx.result)
    }
    return nil
}
```

## 🚀 Running Tests

```bash
# Run all features
go test

# Run specific feature
godog features/addition.feature

# Run with tags
godog --tags=@smoke

# Different formats
godog --format=pretty
godog --format=json
```

## ✅ Pros

- ✅ Industry-standard Gherkin syntax
- ✅ Non-technical stakeholders can read tests
- ✅ Great for acceptance testing
- ✅ Reusable step definitions
- ✅ Integration with CI/CD tools

## ❌ Cons

- ❌ Not ideal for unit tests
- ❌ Requires step definition mapping
- ❌ Can be verbose for simple scenarios
- ❌ Learning curve for Gherkin

## 🔗 Resources

- [Official Documentation](https://github.com/cucumber/godog)
- [Cucumber Documentation](https://cucumber.io/docs/gherkin/)
- [Gherkin Reference](https://cucumber.io/docs/gherkin/reference/)

## 💡 Best Practices

1. **Keep scenarios focused**: One behavior per scenario
2. **Reuse step definitions**: Don't duplicate steps
3. **Use Background**: For common setup
4. **Tag scenarios**: Organize by @smoke, @integration, etc.
5. **Make steps readable**: Write for humans, not machines

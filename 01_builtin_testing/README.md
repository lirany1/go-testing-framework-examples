# Built-in Testing Package

Go's standard `testing` package provides the foundation for all testing in Go. It's part of the standard library, requires no external dependencies, and is the basis for all other testing frameworks.

## 📦 Installation

No installation needed - it's part of Go's standard library!

## 🎯 Features

- **Unit Testing**: Write test functions with `func TestXxx(t *testing.T)`
- **Benchmarking**: Measure performance with `func BenchmarkXxx(b *testing.B)`
- **Parallel Execution**: Run tests concurrently with `t.Parallel()`
- **Subtests**: Organize tests hierarchically with `t.Run()`
- **Table-Driven Tests**: Test multiple scenarios efficiently
- **Test Coverage**: Built-in coverage analysis with `-cover` flag

## 📖 Usage

### Basic Test

```go
func TestSum(t *testing.T) {
    result := Sum(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

### Table-Driven Test

```go
func TestSum(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -2, -3},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Sum(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Sum(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

## 🚀 Running Tests

```bash
# Run all tests
go test

# Run with verbose output
go test -v

# Run with coverage
go test -cover

# Run specific test
go test -run TestSum

# Run benchmarks
go test -bench=.
```

## 📊 Example Output

```
=== RUN   TestSum
=== RUN   TestSum/positive_numbers
=== RUN   TestSum/negative_numbers
=== RUN   TestSum/with_zero
--- PASS: TestSum (0.00s)
    --- PASS: TestSum/positive_numbers (0.00s)
    --- PASS: TestSum/negative_numbers (0.00s)
    --- PASS: TestSum/with_zero (0.00s)
PASS
ok      github.com/lirany1/go-testing-framework-examples/01_builtin_testing    0.002s
```

## ✅ Pros

- ✅ Part of Go standard library (no dependencies)
- ✅ Fast and lightweight
- ✅ Simple and straightforward
- ✅ Excellent IDE support
- ✅ Built-in benchmarking and profiling
- ✅ Native parallel test execution

## ❌ Cons

- ❌ No built-in assertion library (manual `if` checks)
- ❌ Not BDD-style
- ❌ Verbose for complex assertions
- ❌ No built-in mocking (need external tools)

## 🔗 Resources

- [Official Documentation](https://pkg.go.dev/testing)
- [Go Testing Tutorial](https://go.dev/doc/tutorial/add-a-test)
- [Table-Driven Tests in Go](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

## 💡 Best Practices

1. **Use Table-Driven Tests**: Test multiple scenarios efficiently
2. **Name Tests Descriptively**: Use `TestFunctionName_Scenario` pattern
3. **Test One Thing**: Each test should verify a single behavior
4. **Use Subtests**: Organize related tests with `t.Run()`
5. **Check Coverage**: Aim for meaningful coverage, not 100%

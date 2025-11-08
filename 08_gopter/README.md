# Gopter

Gopter is a property-based testing library for Go, inspired by QuickCheck and ScalaCheck. It generates random test inputs to verify properties hold for all possible values.

## 📦 Installation

```bash
go get github.com/leanovate/gopter
```

## 🎯 Features

- **Property-Based Testing**: Verify properties instead of specific examples
- **Random Input Generation**: Automatic test case generation
- **Shrinking**: Minimizes failing inputs to simplest case
- **Generators**: Built-in generators for common types
- **Combinators**: Compose generators for complex types
- **Labels**: Classify and collect statistics about tests

## 📖 Usage

### Basic Property Test

```go
properties := gopter.NewProperties(nil)

properties.Property("addition is commutative", prop.ForAll(
    func(a, b int) bool {
        return a + b == b + a
    },
    gen.Int(), gen.Int(),
))

properties.TestingRun(t)
```

## 🚀 Running Tests

```bash
go test
go test -v  # Verbose output with shrinking details
```

## ✅ Pros

- ✅ Finds edge cases automatically
- ✅ Tests properties, not examples
- ✅ Shrinking reveals minimal failing case
- ✅ Great for algorithms and math
- ✅ Catches unexpected bugs

## ❌ Cons

- ❌ Different thinking paradigm
- ❌ Not suitable for UI/integration tests
- ❌ Can be slow (generates many inputs)
- ❌ Requires defining properties

## 🔗 Resources

- [GitHub Repository](https://github.com/leanovate/gopter)
- [Documentation](https://pkg.go.dev/github.com/leanovate/gopter)

## 💡 Best Practices

1. **Think in properties**: Focus on invariants, not examples
2. **Use appropriate generators**: Match input domain
3. **Set reasonable test counts**: Balance speed vs coverage
4. **Label test cases**: Classify inputs for statistics
5. **Test pure functions**: Property testing works best with deterministic code

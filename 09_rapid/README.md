# Rapid

Rapid is an advanced property-based and model-based testing library for Go. It supports stateful testing and can verify complex state machines and sequences of operations.

## 📦 Installation

```bash
go get pgregory.net/rapid
```

## 🎯 Features

- **Property-Based Testing**: Like Gopter, generates random inputs
- **Model-Based Testing**: Test stateful systems with state machines
- **Shrinking**: Minimizes failing test cases
- **State Machines**: Model system states and transitions
- **Assertions**: Built-in assertion helpers
- **Reproducible**: Tests can be replayed with specific seeds

## 📖 Usage

### Basic Property Test

```go
func TestProperty(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        a := rapid.Int().Draw(t, "a")
        b := rapid.Int().Draw(t, "b")
        
        if a + b != b + a {
            t.Fatalf("commutativity failed")
        }
    })
}
```

### Stateful Testing

```go
rapid.Check(t, func(t *rapid.T) {
    stack := []int{}
    
    for i := 0; i < 100; i++ {
        if rapid.Bool().Draw(t, "push?") {
            stack = append(stack, rapid.Int().Draw(t, "val"))
        } else if len(stack) > 0 {
            stack = stack[:len(stack)-1]
        }
    }
})
```

## 🚀 Running Tests

```bash
go test
go test -v  # See generated values and shrinking
go test -rapid.checks=1000  # More iterations
```

## ✅ Pros

- ✅ Powerful stateful testing
- ✅ Excellent shrinking algorithm
- ✅ Model complex state transitions
- ✅ Good for finding race conditions
- ✅ Reproducible test runs

## ❌ Cons

- ❌ Steep learning curve
- ❌ Complex setup for stateful tests
- ❌ Not suitable for simple unit tests
- ❌ Requires deep system understanding

## 🔗 Resources

- [GitHub Repository](https://github.com/flyingmutant/rapid)
- [Documentation](https://pkg.go.dev/pgregory.net/rapid)

## 💡 Best Practices

1. **Start simple**: Begin with property tests before stateful
2. **Model carefully**: Accurate state models are crucial
3. **Use labels**: Name drawn values for better debugging
4. **Test invariants**: Verify properties hold at every state
5. **Limit iterations**: Balance thoroughness with speed

# Testify

Testify is the most popular testing toolkit for Go, providing rich assertions and mocking capabilities. It extends Go's built-in `testing` package with readable assertions and a powerful mocking framework.

## 📦 Installation

```bash
go get github.com/stretchr/testify
```

## 🎯 Features

- **Rich Assertions**: `assert` and `require` packages with dozens of assertion methods
- **Mocking**: Built-in mock objects with expectation setting
- **Test Suites**: Organize tests into suites with setup/teardown
- **HTTP Testing**: Helpers for testing HTTP handlers

### Difference: assert vs require

- **assert**: Marks test as failed but continues execution
- **require**: Marks test as failed and stops execution immediately

## 📖 Usage

### Basic Assertions

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
    result := Sum(2, 3)
    assert.Equal(t, 5, result, "Sum should equal 5")
}
```

### Using require (fails fast)

```go
func TestDivide(t *testing.T) {
    result := Divide(6, 2)
    require.Equal(t, 3, result)
    // If above fails, execution stops here
}
```

### Mocking

```go
type MockDB struct {
    mock.Mock
}

func (m *MockDB) GetUser(id int) string {
    args := m.Called(id)
    return args.String(0)
}

func TestUserService(t *testing.T) {
    mockDB := new(MockDB)
    mockDB.On("GetUser", 1).Return("John")
    
    result := mockDB.GetUser(1)
    assert.Equal(t, "John", result)
    mockDB.AssertExpectations(t)
}
```

## 🚀 Running Tests

```bash
# Run all tests
go test

# Run with verbose output
go test -v

# Run specific test
go test -run TestSum
```

## 📊 Common Assertions

```go
assert.Equal(t, expected, actual)
assert.NotEqual(t, expected, actual)
assert.Nil(t, object)
assert.NotNil(t, object)
assert.True(t, value)
assert.False(t, value)
assert.Contains(t, haystack, needle)
assert.Len(t, list, expectedLength)
assert.Error(t, err)
assert.NoError(t, err)
```

## ✅ Pros

- ✅ Readable and expressive assertions
- ✅ Built-in mocking framework
- ✅ Large community and wide adoption
- ✅ Reduces boilerplate code significantly
- ✅ Works seamlessly with Go's testing package

## ❌ Cons

- ❌ Not true BDD (still function-based tests)
- ❌ Adds external dependency
- ❌ Learning curve for mock expectations

## 🔗 Resources

- [Official GitHub](https://github.com/stretchr/testify)
- [Documentation](https://pkg.go.dev/github.com/stretchr/testify)
- [Assertion Reference](https://pkg.go.dev/github.com/stretchr/testify/assert)

## 💡 Best Practices

1. **Use require for critical checks**: Stop execution early if fundamental assumptions fail
2. **Use assert for multiple checks**: Continue testing even if one assertion fails
3. **Add descriptive messages**: Help identify failures quickly
4. **Mock external dependencies**: Keep tests fast and isolated
5. **Assert mock expectations**: Always call `AssertExpectations(t)`

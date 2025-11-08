# GoMock

GoMock is an official mocking framework from Google for generating mock objects from Go interfaces. **Note: This project is deprecated as of 2023**, but it's still widely used in existing codebases.

## ⚠️ Deprecation Notice

GoMock has been officially deprecated. Consider using alternatives like:
- [Mockery](https://github.com/vektra/mockery)
- [Testify Mock](https://github.com/stretchr/testify#mock-package)
- [Moq](https://github.com/matryer/moq)

## 📦 Installation

```bash
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@latest
```

## 🎯 Features

- **Auto-generation**: Generate mocks from interfaces
- **Expectation Setting**: Define expected method calls
- **Argument Matching**: Flexible argument matchers
- **Call Ordering**: Verify call sequences
- **Integration**: Works with Go's testing package

## 📖 Usage

### Generate Mocks

```bash
# From interface in file
mockgen -source=db.go -destination=mock_db.go -package=mocks

# From package
mockgen -destination=mock_db.go package/path InterfaceName
```

### Using Mocks

```go
func TestUserService(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockUserRepository(ctrl)
    mockDB.EXPECT().GetUser(1).Return(&User{Name: "John"}, nil)

    // Use mock in tests
}
```

## 🚀 Running Tests

```bash
# Generate mocks first
go generate ./...

# Then run tests
go test
```

## ✅ Pros

- ✅ Official Google project
- ✅ Automatic mock generation
- ✅ Type-safe mocks
- ✅ Comprehensive expectation system

## ❌ Cons

- ❌ **Deprecated since 2023**
- ❌ More verbose than alternatives
- ❌ Requires separate generation step
- ❌ Limited to interfaces only

## 🔗 Resources

- [GitHub Repository](https://github.com/golang/mock)
- [Documentation](https://pkg.go.dev/github.com/golang/mock/gomock)

## 💡 Best Practices

1. **Use go:generate**: Automate mock generation
2. **Consider alternatives**: Mockery or Testify for new projects
3. **Verify expectations**: Always call `ctrl.Finish()`
4. **Use argument matchers**: For flexible expectations

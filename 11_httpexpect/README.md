# httpexpect

httpexpect is a concise and chainable HTTP/REST API testing library for Go. It provides expressive assertions for testing HTTP handlers and APIs with minimal boilerplate.

## 📦 Installation

```bash
go get github.com/gavv/httpexpect/v2
```

## 🎯 Features

- **Chainable API**: Fluent interface for readable tests
- **Rich Assertions**: JSON, headers, status, cookies, etc.
- **Built-in Matchers**: Equal, Contains, Match (regex), etc.
- **WebSocket Support**: Test WebSocket connections
- **FormData & Multipart**: File uploads and forms
- **Integration**: Works with net/http handlers or live servers

## 📖 Usage

### Basic API Test

```go
e := httpexpect.New(t, "http://localhost:8080")

e.GET("/users/1").
    Expect().
    Status(200).
    JSON().Object().
    ValueEqual("name", "John")
```

### Testing HTTP Handler

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Write([]byte(`{"message": "Hello"}`))
})

e := httpexpect.WithConfig(httpexpect.Config{
    Client:   &http.Client{},
    BaseURL:  "http://example.com",
    Reporter: httpexpect.NewAssertReporter(t),
    Printers: []httpexpect.Printer{
        httpexpect.NewDebugPrinter(t, true),
    },
})

e.GET("/").WithHandler(handler).
    Expect().
    Status(200).
    JSON().Object().
    ValueEqual("message", "Hello")
```

## 🚀 Running Tests

```bash
go test
go test -v  # See HTTP requests/responses
```

## ✅ Pros

- ✅ Extremely readable test code
- ✅ Comprehensive HTTP assertions
- ✅ Works with handlers or live servers
- ✅ Excellent for API testing
- ✅ JSON path navigation
- ✅ Schema validation support

## ❌ Cons

- ❌ HTTP-only (not for general testing)
- ❌ Learning curve for chainable API
- ❌ Can be verbose for simple tests
- ❌ Focused on REST/JSON APIs

## 🔗 Resources

- [GitHub Repository](https://github.com/gavv/httpexpect)
- [Documentation](https://pkg.go.dev/github.com/gavv/httpexpect/v2)
- [Examples](https://github.com/gavv/httpexpect/tree/master/_examples)

## 💡 Best Practices

1. **Chain assertions**: Make tests readable
2. **Test handlers directly**: Faster than live servers
3. **Use WithHandler**: For unit testing handlers
4. **Validate JSON schema**: Ensure API contracts
5. **Check all response aspects**: Status, headers, body

# GoConvey

GoConvey is a BDD testing framework for Go with a unique feature: a beautiful Web UI for viewing test results in real-time. It provides readable test syntax and live reload capabilities.

## 📦 Installation

```bash
go get github.com/smartystreets/goconvey
```

## 🎯 Features

- **BDD Style**: Convey/So syntax for readable tests
- **Web UI**: Beautiful browser-based test runner with live reload
- **Nested Contexts**: Natural test organization
- **Rich Assertions**: Many built-in matchers (ShouldEqual, ShouldBeNil, etc.)
- **Auto-Watch**: Automatically reruns tests on file changes
- **Coverage Visualization**: See code coverage in the UI

## 📖 Usage

### Basic Structure

```go
Convey("Sum function", t, func() {
    Convey("should add two numbers", func() {
        So(Sum(2, 3), ShouldEqual, 5)
    })
})
```

### Nested Contexts

```go
Convey("Calculator", t, func() {
    Convey("Addition", func() {
        Convey("with positive numbers", func() {
            So(Sum(2, 3), ShouldEqual, 5)
        })
    })
})
```

## 🚀 Running Tests

```bash
# Standard go test
go test

# With GoConvey Web UI (recommended)
goconvey

# Then open browser to http://localhost:8080
```

## 🌐 Web UI Features

The GoConvey Web UI provides:
- Real-time test results
- Auto-rerun on file save
- Code coverage visualization
- Test history
- Notifications on failures

## 📊 Common Assertions

```go
So(actual, ShouldEqual, expected)
So(actual, ShouldNotEqual, expected)
So(object, ShouldBeNil)
So(object, ShouldNotBeNil)
So(slice, ShouldContain, element)
So(slice, ShouldHaveLength, 3)
So(err, ShouldBeError)
So(value, ShouldBeTrue)
```

## ✅ Pros

- ✅ Beautiful, intuitive Web UI
- ✅ Live reload and watch mode
- ✅ Readable BDD-style syntax
- ✅ Visual code coverage
- ✅ Good for TDD workflow
- ✅ Nested contexts for organization

## ❌ Cons

- ❌ Less actively maintained (as of 2025)
- ❌ Web UI adds overhead
- ❌ Smaller community compared to alternatives
- ❌ Not as feature-rich as Ginkgo

## 🔗 Resources

- [Official Website](https://smartystreets.github.io/goconvey/)
- [GitHub Repository](https://github.com/smartystreets/goconvey)
- [Documentation](https://github.com/smartystreets/goconvey/wiki)

## 💡 Best Practices

1. **Use the Web UI**: It's GoConvey's killer feature
2. **Nest Convey blocks**: Organize tests logically
3. **Use descriptive strings**: Make the UI output readable
4. **One assertion per Convey**: Keep tests focused
5. **Enable watch mode**: Get instant feedback during development

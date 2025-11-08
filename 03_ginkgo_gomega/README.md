# Ginkgo + Gomega

Ginkgo is a BDD-style testing framework for Go, paired with Gomega as its matcher/assertion library. Together they provide expressive, readable tests using `Describe`, `Context`, and `It` blocks.

## 📦 Installation

```bash
# Install Ginkgo and Gomega
go get github.com/onsi/ginkgo/v2
go get github.com/onsi/gomega

# Install Ginkgo CLI (recommended)
go install github.com/onsi/ginkgo/v2/ginkgo@latest
```

## 🎯 Features

- **BDD Style**: Write tests in Describe/Context/It structure
- **Rich Matchers**: Gomega provides expressive assertions
- **Parallel Execution**: Built-in support for concurrent tests
- **Test Organization**: Logical grouping with nested contexts
- **Setup/Teardown**: BeforeEach, AfterEach, BeforeSuite, AfterSuite hooks
- **Focus/Pending**: FDescribe, FIt for focused tests, PIt for pending

## 📖 Usage

### Basic Structure

```go
var _ = Describe("Calculator", func() {
    Context("when adding numbers", func() {
        It("should return the correct sum", func() {
            result := Sum(2, 3)
            Expect(result).To(Equal(5))
        })
    })
})
```

### With Setup/Teardown

```go
var _ = Describe("Database", func() {
    var db *Database

    BeforeEach(func() {
        db = NewDatabase()
    })

    AfterEach(func() {
        db.Close()
    })

    It("should connect successfully", func() {
        Expect(db.IsConnected()).To(BeTrue())
    })
})
```

## 🚀 Running Tests

```bash
# Using Ginkgo CLI (recommended)
ginkgo

# With verbose output
ginkgo -v

# Run specific suite
ginkgo ./03_ginkgo_gomega

# Run in parallel
ginkgo -p

# Using go test
go test
```

## 📊 Common Matchers

```go
Expect(value).To(Equal(expected))
Expect(value).To(BeNumerically(">", 5))
Expect(slice).To(ContainElement(item))
Expect(slice).To(HaveLen(3))
Expect(str).To(ContainSubstring("hello"))
Expect(err).To(HaveOccurred())
Expect(err).NotTo(HaveOccurred())
Expect(value).To(BeNil())
Expect(value).To(BeTrue())
```

## ✅ Pros

- ✅ Clean, readable BDD-style syntax
- ✅ Excellent test organization with nested contexts
- ✅ Rich ecosystem of matchers
- ✅ Built-in parallel test execution
- ✅ Great CLI with watch mode and focus features
- ✅ Comprehensive hooks for setup/teardown

## ❌ Cons

- ❌ Steeper learning curve than standard testing
- ❌ Additional dependencies
- ❌ Different paradigm from Go's standard testing
- ❌ Can be verbose for simple tests

## 🔗 Resources

- [Ginkgo Documentation](https://onsi.github.io/ginkgo/)
- [Gomega Documentation](https://onsi.github.io/gomega/)
- [GitHub - Ginkgo](https://github.com/onsi/ginkgo)
- [GitHub - Gomega](https://github.com/onsi/gomega)

## 💡 Best Practices

1. **Use Describe for components**: Group related tests together
2. **Use Context for scenarios**: Different conditions or states
3. **Use It for expectations**: Single behavior per It block
4. **Leverage BeforeEach**: Set up common test state
5. **Focus during development**: Use FIt or FDescribe to run specific tests
6. **Keep It blocks small**: One logical assertion per It
7. **Use descriptive strings**: Make test output readable

# Go Testing Frameworks Examples

> **Version 1.0.0** | A curated collection of Go testing framework examples — from built-in unit tests to advanced BDD, property-based, and integration testing.

![Go Version](https://img.shields.io/badge/go-1.23+-blue.svg)
![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![Build Status](https://github.com/lirany1/go-testing-framework-examples/actions/workflows/go.yml/badge.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![GitHub Stars](https://img.shields.io/github/stars/lirany1/go-testing-framework-examples?style=social)
[![Buy Me A Coffee](https://img.shields.io/badge/Buy%20Me%20A%20Coffee-support-yellow.svg)](https://buymeacoffee.com/liran80v)

## 📋 About

**Go Testing Frameworks Examples** is a comprehensive educational repository showcasing 11 popular testing frameworks in the Go ecosystem. Each framework is demonstrated with production-ready, well-documented code examples covering various testing paradigms:

- 🧪 **Unit Testing** - Native Go testing and Testify
- 🎭 **BDD Testing** - Ginkgo, GoConvey, Godog, Gauge
- 🎯 **Mocking** - GoMock for interface mocking
- 🔀 **Property-Based** - Gopter for generative testing
- 🤖 **Model-Based** - Rapid for stateful testing
- 🐳 **Integration** - Testcontainers for Docker-based tests
- 🌐 **API Testing** - httpexpect for REST endpoints

### 🏷️ Tags

`go` `golang` `testing` `bdd` `tdd` `unit-testing` `integration-testing` `property-based-testing` `mocking` `testify` `ginkgo` `gomega` `goconvey` `gomock` `godog` `cucumber` `gauge` `gopter` `rapid` `testcontainers` `httpexpect` `testing-frameworks` `go-testing` `test-examples` `best-practices` `tutorial` `learning` `education`

## 📖 Overview

This repository contains practical, runnable examples for the most popular testing frameworks and tools in the Go ecosystem. Each framework is demonstrated with clean, well-documented code that showcases real-world testing scenarios.

Whether you're new to Go testing or looking to explore advanced testing techniques, this repository provides hands-on examples for:
- Unit testing with built-in `testing` package
- Assertion libraries and mocking frameworks
- BDD (Behavior-Driven Development) approaches
- Property-based and model-based testing
- Integration testing with real services
- API/HTTP testing

## 🚀 Quick Start

### Prerequisites
- Go 1.23 or higher
- Docker (required for Testcontainers examples)

### Installation

```bash
# Clone the repository
git clone https://github.com/lirany1/go-testing-framework-examples.git

# Navigate to the project directory
cd go-testing-framework-examples

# Download dependencies
go mod tidy

# Run all tests (except Gauge and Testcontainers)
go test ./01_builtin_testing/... ./02_testify/... ./03_ginkgo_gomega/... \
        ./04_goconvey/... ./05_gomock/... ./06_godog/... \
        ./08_gopter/... ./09_rapid/... ./11_httpexpect/...
```

**Note:** 
- **Gauge** (07_gauge) requires the Gauge CLI to run tests. See [Gauge documentation](https://docs.gauge.org/).
- **Testcontainers** (10_testcontainers_go) requires Docker to be running for integration tests.

### Run Tests for a Specific Framework

```bash
# Built-in testing
go test ./01_builtin_testing/...

# Testify
go test ./02_testify/...

# Ginkgo
go test ./03_ginkgo_gomega/...

# GoConvey
go test ./04_goconvey/...

# And so on...
```

## 📊 Framework Comparison

| # | Framework | Type | Ideal for | Key Strength | Limitation |
|---|-----------|------|-----------|--------------|------------|
| 1 | [Built-in testing](./01_builtin_testing) | Unit | Core Go tests | Fast & native | No advanced assertions |
| 2 | [Testify](./02_testify) | Modular | Unit / Integration | Readable assertions, mocks | Not BDD |
| 3 | [Ginkgo + Gomega](./03_ginkgo_gomega) | BDD | Behavior-driven | Clean DSL, strong ecosystem | Heavy dependencies |
| 4 | [GoConvey](./04_goconvey) | BDD | Unit / Acceptance | Web UI, live reload | Less maintained |
| 5 | [GoMock](./05_gomock) | Mocking | Unit | Auto-generated mocks | Deprecated |
| 6 | [Godog](./06_godog) | BDD (Cucumber) | Acceptance | Gherkin syntax | Requires step mapping |
| 7 | [Gauge](./07_gauge) | BDD / Keyword | Acceptance | Natural language specs | Requires setup |
| 8 | [Gopter](./08_gopter) | Property-based | Algorithms | Randomized tests | Not for API/UI |
| 9 | [Rapid](./09_rapid) | Model-based | Stateful systems | Transitions & states | Complex setup |
| 10 | [Testcontainers-go](./10_testcontainers_go) | Integration | Distributed systems | Real services | Requires Docker |
| 11 | [httpexpect](./11_httpexpect) | API | REST validation | Chainable, expressive | HTTP-only |

## 📂 Repository Structure

```
go-testing-framework-examples/
│
├── 01_builtin_testing/       # Go's native testing package
├── 02_testify/               # Assertions and mocking with Testify
├── 03_ginkgo_gomega/         # BDD testing with Ginkgo and Gomega
├── 04_goconvey/              # BDD with Web UI
├── 05_gomock/                # Mock generation for interfaces
├── 06_godog/                 # Cucumber-style BDD with Gherkin
├── 07_gauge/                 # Acceptance testing with natural language
├── 08_gopter/                # Property-based testing
├── 09_rapid/                 # Model-based testing for stateful systems
├── 10_testcontainers_go/    # Integration testing with Docker containers
├── 11_httpexpect/            # HTTP/API testing
└── .github/workflows/        # CI/CD pipeline
```

Each directory contains:
- Working code examples
- README with installation and usage instructions
- Comments explaining key concepts

## 🧪 What's Included

### 1. [Built-in testing](./01_builtin_testing)
Go's standard testing package - the foundation of all Go testing.

### 2. [Testify](./02_testify)
The most popular assertion library with built-in mocking support.

### 3. [Ginkgo + Gomega](./03_ginkgo_gomega)
BDD-style testing with descriptive test structure (Describe/Context/It).

### 4. [GoConvey](./04_goconvey)
BDD framework with a beautiful Web UI for real-time test results.

### 5. [GoMock](./05_gomock)
Automatic mock generation from interfaces (note: deprecated but widely used).

### 6. [Godog](./06_godog)
Cucumber implementation for Go - write tests in plain English with Gherkin.

### 7. [Gauge](./07_gauge)
Keyword-driven acceptance testing from ThoughtWorks.

### 8. [Gopter](./08_gopter)
Property-based testing - validate properties with random inputs.

### 9. [Rapid](./09_rapid)
Advanced model-based testing for stateful systems.

### 10. [Testcontainers-go](./10_testcontainers_go)
Integration testing with real Docker containers (databases, message queues, etc.).

### 11. [httpexpect](./11_httpexpect)
Elegant HTTP API testing with chainable assertions.

## 🛠️ Running Tests

### All Tests
```bash
go test ./...
```

### With Verbose Output
```bash
go test -v ./...
```

### Specific Framework
```bash
go test ./01_builtin_testing/...
```

### With Coverage
```bash
go test -cover ./...
```

## 📚 Learning Path

If you're new to Go testing, we recommend following this order:

1. **Start with Built-in testing** - Understand Go's native testing approach
2. **Add Testify** - Learn assertion libraries and mocking
3. **Explore Ginkgo** - Understand BDD and descriptive test structures
4. **Try httpexpect** - Learn API testing patterns
5. **Experiment with Property-based testing** - Explore advanced testing techniques
6. **Use Testcontainers** - Master integration testing with real services

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Add new framework examples
- Improve existing examples
- Fix bugs or typos
- Enhance documentation

Please ensure:
- Code follows Go conventions (`gofmt`, `golint`)
- All examples compile and run successfully
- Each example includes clear documentation

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ☕ Support the Author

If you found this project useful, you can support me here:

[![Buy Me A Coffee](https://img.shields.io/badge/Buy%20Me%20A%20Coffee-support-yellow.svg?style=for-the-badge&logo=buy-me-a-coffee)](https://buymeacoffee.com/liran80v)

Your support helps maintain and expand this educational resource!

## 📞 Contact

- GitHub: [@lirany1](https://github.com/lirany1)
- Repository: [go-testing-framework-examples](https://github.com/lirany1/go-testing-framework-examples)

## 🌟 Acknowledgments

Special thanks to the maintainers and contributors of all the testing frameworks featured in this repository. The Go testing ecosystem is rich and diverse thanks to their efforts.

---

**Made with ❤️ for the Go Testing community**

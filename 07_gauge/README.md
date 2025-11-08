# Gauge

Gauge is an acceptance testing framework from ThoughtWorks that uses Markdown specifications and supports keyword-driven testing. Tests are written in natural language with reusable steps.

## 📦 Installation

```bash
# Install Gauge
brew install gauge  # macOS
choco install gauge # Windows
# Or download from https://gauge.org

# Install Go plugin
gauge install go
```

## 🎯 Features

- **Markdown Specs**: Write tests in Markdown format
- **Keyword-Driven**: Reusable step implementations
- **Multi-Language**: Supports multiple programming languages
- **Rich Reports**: HTML reports with screenshots
- **Data-Driven**: Run same scenario with different data
- **Parallel Execution**: Run specs in parallel

## 📖 Usage

### Specification File (.spec)

```markdown
# Login Flow

## Valid login
* Navigate to the login page
* Enter username "admin" and password "admin123"
* Click login
* Verify that dashboard is visible
```

### Step Implementation (Go)

```go
func NavigateToLoginPage() {
    // Implementation
}

func EnterCredentials(username, password string) {
    // Implementation
}
```

## 🚀 Running Tests

```bash
# Initialize Gauge project
gauge init go

# Run all specs
gauge run specs

# Run specific spec
gauge run specs/login.spec

# Run with tags
gauge run --tags "smoke" specs

# Parallel execution
gauge run --parallel specs
```

## ✅ Pros

- ✅ Natural language specifications
- ✅ Multi-language support
- ✅ Excellent for acceptance testing
- ✅ Rich HTML reports
- ✅ Reusable step library
- ✅ Good IDE support (VS Code plugin)

## ❌ Cons

- ❌ Not suitable for unit tests
- ❌ Requires Gauge installation
- ❌ Additional setup complexity
- ❌ Smaller community than Cucumber

## 🔗 Resources

- [Official Website](https://gauge.org)
- [Documentation](https://docs.gauge.org)
- [GitHub](https://github.com/getgauge/gauge)
- [Go Plugin](https://github.com/getgauge-contrib/gauge-go)

## 💡 Best Practices

1. **Keep specs focused**: One feature per spec file
2. **Reuse steps**: Build a library of common steps
3. **Use concepts**: Group related steps together
4. **Tag scenarios**: Organize with @smoke, @regression, etc.
5. **Parameterize steps**: Make steps flexible with arguments

# Testcontainers-go

Testcontainers is a library for integration testing with real Docker containers. Spin up databases, message queues, or any Docker service for realistic integration tests.

## 📦 Installation

```bash
go get github.com/testcontainers/testcontainers-go
```

## 🎯 Features

- **Real Services**: Test against actual databases, Redis, Kafka, etc.
- **Isolated Environment**: Each test gets fresh containers
- **Auto-Cleanup**: Containers removed after tests
- **Wait Strategies**: Wait for services to be ready
- **Network Control**: Custom networks and port mapping
- **Docker Compose**: Support for docker-compose files

## 📖 Usage

### Basic Container

```go
ctx := context.Background()
req := testcontainers.ContainerRequest{
    Image:        "redis:latest",
    ExposedPorts: []string{"6379/tcp"},
    WaitingFor:   wait.ForLog("Ready to accept connections"),
}
container, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
    ContainerRequest: req,
    Started:          true,
})
defer container.Terminate(ctx)
```

## 🚀 Running Tests

```bash
# Requires Docker to be running
docker --version

# Run tests
go test
```

## ✅ Pros

- ✅ Test with real services (not mocks)
- ✅ Isolated test environment
- ✅ Reproducible across machines
- ✅ Great for integration/E2E tests
- ✅ Supports many pre-built containers

## ❌ Cons

- ❌ Requires Docker installation
- ❌ Slower than unit tests
- ❌ Higher resource usage
- ❌ Not suitable for CI without Docker

## 🔗 Resources

- [Official Documentation](https://golang.testcontainers.org/)
- [GitHub Repository](https://github.com/testcontainers/testcontainers-go)
- [Examples](https://golang.testcontainers.org/examples/)

## 💡 Best Practices

1. **Always defer Terminate()**: Clean up containers
2. **Use wait strategies**: Ensure services are ready
3. **Reuse containers**: Within test suites when possible
4. **Tag tests**: Use build tags for integration tests
5. **Check Docker**: Verify Docker is running before tests

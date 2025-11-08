package testcontainers

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestRedisContainer demonstrates using Testcontainers with Redis.
func TestRedisContainer(t *testing.T) {
	// Skip if Docker is not available
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()

	// Define container request
	req := testcontainers.ContainerRequest{
		Image:        "redis:7-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	// Start Redis container
	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start Redis container: %v", err)
	}

	// Always clean up the container
	defer func() {
		if err := redisContainer.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	}()

	// Get the container host and port
	host, err := redisContainer.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}

	port, err := redisContainer.MappedPort(ctx, "6379")
	if err != nil {
		t.Fatalf("Failed to get container port: %v", err)
	}

	// Connect to Redis
	redisAddr := fmt.Sprintf("%s:%s", host, port.Port())
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	defer client.Close()

	// Test Redis operations
	t.Run("SET and GET operations", func(t *testing.T) {
		key := "test-key"
		value := "test-value"

		// Set a key
		err := client.Set(ctx, key, value, 0).Err()
		if err != nil {
			t.Fatalf("Failed to SET key: %v", err)
		}

		// Get the key
		result, err := client.Get(ctx, key).Result()
		if err != nil {
			t.Fatalf("Failed to GET key: %v", err)
		}

		if result != value {
			t.Errorf("Expected %s, got %s", value, result)
		}
	})

	t.Run("Key expiration", func(t *testing.T) {
		key := "expiring-key"
		value := "will-expire"
		expiration := 1 * time.Second

		// Set key with expiration
		err := client.Set(ctx, key, value, expiration).Err()
		if err != nil {
			t.Fatalf("Failed to SET key with expiration: %v", err)
		}

		// Verify key exists
		exists, err := client.Exists(ctx, key).Result()
		if err != nil {
			t.Fatalf("Failed to check if key exists: %v", err)
		}
		if exists != 1 {
			t.Error("Key should exist immediately after setting")
		}

		// Wait for expiration
		time.Sleep(2 * time.Second)

		// Verify key expired
		exists, err = client.Exists(ctx, key).Result()
		if err != nil {
			t.Fatalf("Failed to check if key exists: %v", err)
		}
		if exists != 0 {
			t.Error("Key should have expired")
		}
	})

	t.Run("List operations", func(t *testing.T) {
		listKey := "test-list"

		// Push values to list
		err := client.RPush(ctx, listKey, "item1", "item2", "item3").Err()
		if err != nil {
			t.Fatalf("Failed to RPUSH: %v", err)
		}

		// Get list length
		length, err := client.LLen(ctx, listKey).Result()
		if err != nil {
			t.Fatalf("Failed to LLEN: %v", err)
		}
		if length != 3 {
			t.Errorf("Expected list length 3, got %d", length)
		}

		// Get all items
		items, err := client.LRange(ctx, listKey, 0, -1).Result()
		if err != nil {
			t.Fatalf("Failed to LRANGE: %v", err)
		}

		expected := []string{"item1", "item2", "item3"}
		if len(items) != len(expected) {
			t.Errorf("Expected %d items, got %d", len(expected), len(items))
		}

		for i, item := range items {
			if item != expected[i] {
				t.Errorf("At index %d: expected %s, got %s", i, expected[i], item)
			}
		}
	})

	t.Run("Hash operations", func(t *testing.T) {
		hashKey := "test-hash"

		// Set hash fields
		err := client.HSet(ctx, hashKey, map[string]interface{}{
			"name":  "John Doe",
			"age":   "30",
			"email": "john@example.com",
		}).Err()
		if err != nil {
			t.Fatalf("Failed to HSET: %v", err)
		}

		// Get single field
		name, err := client.HGet(ctx, hashKey, "name").Result()
		if err != nil {
			t.Fatalf("Failed to HGET: %v", err)
		}
		if name != "John Doe" {
			t.Errorf("Expected 'John Doe', got '%s'", name)
		}

		// Get all fields
		allFields, err := client.HGetAll(ctx, hashKey).Result()
		if err != nil {
			t.Fatalf("Failed to HGETALL: %v", err)
		}
		if len(allFields) != 3 {
			t.Errorf("Expected 3 fields, got %d", len(allFields))
		}
	})

	t.Run("Set operations", func(t *testing.T) {
		setKey := "test-set"

		// Add members to set
		err := client.SAdd(ctx, setKey, "member1", "member2", "member3").Err()
		if err != nil {
			t.Fatalf("Failed to SADD: %v", err)
		}

		// Check if member exists
		exists, err := client.SIsMember(ctx, setKey, "member2").Result()
		if err != nil {
			t.Fatalf("Failed to SISMEMBER: %v", err)
		}
		if !exists {
			t.Error("member2 should exist in set")
		}

		// Get all members
		members, err := client.SMembers(ctx, setKey).Result()
		if err != nil {
			t.Fatalf("Failed to SMEMBERS: %v", err)
		}
		if len(members) != 3 {
			t.Errorf("Expected 3 members, got %d", len(members))
		}
	})

	t.Run("Atomic operations", func(t *testing.T) {
		counterKey := "counter"

		// Increment counter
		newValue, err := client.Incr(ctx, counterKey).Result()
		if err != nil {
			t.Fatalf("Failed to INCR: %v", err)
		}
		if newValue != 1 {
			t.Errorf("Expected 1, got %d", newValue)
		}

		// Increment by specific amount
		newValue, err = client.IncrBy(ctx, counterKey, 5).Result()
		if err != nil {
			t.Fatalf("Failed to INCRBY: %v", err)
		}
		if newValue != 6 {
			t.Errorf("Expected 6, got %d", newValue)
		}

		// Decrement
		newValue, err = client.Decr(ctx, counterKey).Result()
		if err != nil {
			t.Fatalf("Failed to DECR: %v", err)
		}
		if newValue != 5 {
			t.Errorf("Expected 5, got %d", newValue)
		}
	})
}

// TestRedisContainer_MultipleOperations demonstrates complex scenarios.
func TestRedisContainer_MultipleOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()

	// Start Redis container
	req := testcontainers.ContainerRequest{
		Image:        "redis:7-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start Redis container: %v", err)
	}
	defer func() {
		_ = redisContainer.Terminate(ctx) // Error ignored in defer - container cleanup
	}()

	// Get connection details
	host, _ := redisContainer.Host(ctx)
	port, _ := redisContainer.MappedPort(ctx, "6379")

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port.Port()),
	})
	defer client.Close()

	// Simulate a user session management scenario
	t.Run("User session management", func(t *testing.T) {
		userID := "user:123"
		sessionData := map[string]interface{}{
			"username":  "johndoe",
			"email":     "john@example.com",
			"loginTime": time.Now().Unix(),
			"isActive":  "true",
		}

		// Store session
		err := client.HSet(ctx, userID, sessionData).Err()
		if err != nil {
			t.Fatalf("Failed to store session: %v", err)
		}

		// Set session expiration (30 minutes)
		err = client.Expire(ctx, userID, 30*time.Minute).Err()
		if err != nil {
			t.Fatalf("Failed to set expiration: %v", err)
		}

		// Retrieve session
		session, err := client.HGetAll(ctx, userID).Result()
		if err != nil {
			t.Fatalf("Failed to retrieve session: %v", err)
		}

		if session["username"] != "johndoe" {
			t.Errorf("Username mismatch")
		}

		// Check TTL
		ttl, err := client.TTL(ctx, userID).Result()
		if err != nil {
			t.Fatalf("Failed to get TTL: %v", err)
		}
		if ttl <= 0 {
			t.Error("Session should have positive TTL")
		}
	})
}

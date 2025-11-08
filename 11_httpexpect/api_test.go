package httpexpect

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

// User represents a user in our API.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// createAPIHandler creates a simple REST API handler for testing.
func createAPIHandler() http.Handler {
	mux := http.NewServeMux()

	// GET /health - Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
		})
	})

	// GET /users - List all users
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		users := []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(users)
	})

	// GET /users/{id} - Get specific user
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Simple routing - in real app use a router library
		id := r.URL.Path[len("/users/"):]
		if id == "1" {
			user := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(user)
		} else if id == "999" {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "User not found", http.StatusNotFound)
		}
	})

	// POST /users - Create user
	mux.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		user.ID = 3 // Simulate ID assignment
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(user)
	})

	return mux
}

// TestHealthEndpoint demonstrates testing a simple health check.
func TestHealthEndpoint(t *testing.T) {
	handler := createAPIHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	e.GET("/health").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("status", "healthy")
}

// TestListUsers demonstrates testing a list endpoint.
func TestListUsers(t *testing.T) {
	handler := createAPIHandler()
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	obj := e.GET("/users").
		Expect().
		Status(http.StatusOK).
		ContentType("application/json").
		JSON().Array()

	// Verify array length
	obj.Length().IsEqual(2)

	// Verify first user
	obj.Value(0).Object().
		ValueEqual("id", 1).
		ValueEqual("name", "John Doe").
		ValueEqual("email", "john@example.com")

	// Verify second user
	obj.Value(1).Object().
		ValueEqual("id", 2).
		ValueEqual("name", "Jane Smith")
}

// TestGetUser demonstrates testing a specific resource endpoint.
func TestGetUser(t *testing.T) {
	handler := createAPIHandler()
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	t.Run("Existing user", func(t *testing.T) {
		obj := e.GET("/users/1").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		obj.HasValue("id", 1)
		obj.HasValue("name", "John Doe")
		obj.HasValue("email", "john@example.com")

		// Alternative: check all fields at once
		obj.IsEqual(User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		})
	})

	t.Run("Non-existent user", func(t *testing.T) {
		e.GET("/users/999").
			Expect().
			Status(http.StatusNotFound)
	})
}

// TestCreateUser demonstrates testing POST requests with JSON body.
func TestCreateUser(t *testing.T) {
	handler := createAPIHandler()
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	newUser := map[string]interface{}{
		"name":  "Alice Johnson",
		"email": "alice@example.com",
	}

	obj := e.POST("/users/create").
		WithJSON(newUser).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	obj.HasValue("id", 3)
	obj.HasValue("name", "Alice Johnson")
	obj.HasValue("email", "alice@example.com")

	// Verify specific fields exist
	obj.ContainsKey("id")
	obj.ContainsKey("name")
	obj.ContainsKey("email")
}

// TestWithHandler demonstrates testing handlers directly (without server).
func TestWithHandler(t *testing.T) {
	handler := createAPIHandler()

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	e.GET("/health").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("status", "healthy")
}

// TestJSONAssertions demonstrates various JSON assertion methods.
func TestJSONAssertions(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"name":   "John",
			"age":    30,
			"active": true,
			"tags":   []string{"developer", "golang"},
			"address": map[string]string{
				"city":    "New York",
				"country": "USA",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	obj := e.GET("/").
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	// Test different value types
	obj.HasValue("name", "John")
	obj.HasValue("age", 30)
	obj.HasValue("active", true)

	// Test nested objects
	obj.Value("address").Object().
		ValueEqual("city", "New York").
		ValueEqual("country", "USA")

	// Test arrays
	tags := obj.Value("tags").Array()
	tags.Length().IsEqual(2)
	tags.ContainsOnly("developer", "golang")
	tags.Value(0).IsEqual("developer")

	// Test key existence
	obj.ContainsKey("name")
	obj.ContainsKey("address")
	obj.NotContainsKey("phone")
}

// TestHeadersAndCookies demonstrates testing HTTP headers and cookies.
func TestHeadersAndCookies(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set custom headers
		w.Header().Set("X-Custom-Header", "CustomValue")
		w.Header().Set("X-Request-ID", "12345")

		// Set a cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: "abc123",
			Path:  "/",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK")) // Error ignored - test handler
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	resp := e.GET("/").
		Expect().
		Status(http.StatusOK)

	// Test headers
	resp.Header("X-Custom-Header").IsEqual("CustomValue")
	resp.Header("X-Request-ID").IsEqual("12345")

	// Test cookies
	resp.Cookie("session").Value().IsEqual("abc123")
}

// TestMethodNotAllowed demonstrates testing error responses.
func TestMethodNotAllowed(t *testing.T) {
	handler := createAPIHandler()
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	// GET endpoint doesn't support POST
	e.POST("/users").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

// TestQueryParameters demonstrates testing with query parameters.
func TestQueryParameters(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		filter := query.Get("filter")
		limit := query.Get("limit")

		response := map[string]string{
			"filter": filter,
			"limit":  limit,
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	e.GET("/search").
		WithQuery("filter", "active").
		WithQuery("limit", "10").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("filter", "active").
		ValueEqual("limit", "10")
}

// TestChainedRequests demonstrates making multiple related requests.
func TestChainedRequests(t *testing.T) {
	handler := createAPIHandler()
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	// First, check health
	e.GET("/health").
		Expect().
		Status(http.StatusOK)

	// Then, list users
	e.GET("/users").
		Expect().
		Status(http.StatusOK).
		JSON().Array().
		Length().IsEqual(2)

	// Finally, get specific user
	e.GET("/users/1").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("name", "John Doe")
}

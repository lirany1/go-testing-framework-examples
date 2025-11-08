package gomock

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

// TestUserService_GetUserName demonstrates basic mock usage.
func TestUserService_GetUserName(t *testing.T) {
	// Create a controller to manage mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // Verify all expectations were met

	// Create a mock repository
	mockRepo := NewMockUserRepository(ctrl)

	// Set expectations: when GetUser(1) is called, return a user
	expectedUser := &User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	mockRepo.EXPECT().GetUser(1).Return(expectedUser, nil)

	// Create service with mock
	service := NewUserService(mockRepo)

	// Test the service
	name, err := service.GetUserName(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if name != "John Doe" {
		t.Errorf("Expected 'John Doe', got '%s'", name)
	}
}

// TestUserService_GetUserName_NotFound demonstrates error handling with mocks.
func TestUserService_GetUserName_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Simulate a "user not found" error
	mockRepo.EXPECT().GetUser(999).Return(nil, errors.New("user not found"))

	service := NewUserService(mockRepo)
	name, err := service.GetUserName(999)

	if err == nil {
		t.Error("Expected error, got nil")
	}
	if name != "" {
		t.Errorf("Expected empty name, got '%s'", name)
	}
}

// TestUserService_CreateUser demonstrates mocking SaveUser.
func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Use gomock.Any() to match any User argument
	mockRepo.EXPECT().SaveUser(gomock.Any()).Return(nil)

	service := NewUserService(mockRepo)
	err := service.CreateUser("Alice", "alice@example.com")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// TestUserService_CreateUser_WithMatcher demonstrates custom argument matching.
func TestUserService_CreateUser_WithMatcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Match specific user properties
	mockRepo.EXPECT().SaveUser(gomock.AssignableToTypeOf(&User{})).DoAndReturn(
		func(user *User) error {
			if user.Name != "Bob" {
				t.Errorf("Expected user name 'Bob', got '%s'", user.Name)
			}
			if user.Email != "bob@example.com" {
				t.Errorf("Expected email 'bob@example.com', got '%s'", user.Email)
			}
			return nil
		},
	)

	service := NewUserService(mockRepo)
	err := service.CreateUser("Bob", "bob@example.com")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// TestUserService_RemoveUser demonstrates mocking DeleteUser.
func TestUserService_RemoveUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Expect DeleteUser to be called with ID 1
	mockRepo.EXPECT().DeleteUser(1).Return(nil)

	service := NewUserService(mockRepo)
	err := service.RemoveUser(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// TestUserService_GetAllUserNames demonstrates mocking methods that return slices.
func TestUserService_GetAllUserNames(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Return a list of users
	users := []*User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	}
	mockRepo.EXPECT().ListUsers().Return(users, nil)

	service := NewUserService(mockRepo)
	names, err := service.GetAllUserNames()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedNames := []string{"Alice", "Bob", "Charlie"}
	if len(names) != len(expectedNames) {
		t.Errorf("Expected %d names, got %d", len(expectedNames), len(names))
	}

	for i, name := range names {
		if name != expectedNames[i] {
			t.Errorf("Expected name '%s' at index %d, got '%s'", expectedNames[i], i, name)
		}
	}
}

// TestUserService_CallOrder demonstrates verifying call order.
func TestUserService_CallOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Use InOrder to enforce call sequence
	gomock.InOrder(
		mockRepo.EXPECT().GetUser(1).Return(&User{ID: 1, Name: "Alice"}, nil),
		mockRepo.EXPECT().DeleteUser(1).Return(nil),
	)

	service := NewUserService(mockRepo)

	// Must call in this exact order
	_, _ = service.GetUserName(1)
	_ = service.RemoveUser(1)
}

// TestUserService_Times demonstrates controlling how many times a method is called.
func TestUserService_Times(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Expect GetUser to be called exactly 3 times
	mockRepo.EXPECT().GetUser(gomock.Any()).Return(&User{Name: "Test"}, nil).Times(3)

	service := NewUserService(mockRepo)

	// Call exactly 3 times
	_, _ = service.GetUserName(1)
	_, _ = service.GetUserName(2)
	_, _ = service.GetUserName(3)
}

// TestUserService_AnyTimes demonstrates allowing any number of calls.
func TestUserService_AnyTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Allow GetUser to be called any number of times (including zero)
	mockRepo.EXPECT().GetUser(gomock.Any()).Return(&User{Name: "Test"}, nil).AnyTimes()

	service := NewUserService(mockRepo)

	// Can call 0, 1, or many times
	_, _ = service.GetUserName(1)
}

// TestUserService_DoAndReturn demonstrates custom return logic.
func TestUserService_DoAndReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// Custom logic based on input
	mockRepo.EXPECT().GetUser(gomock.Any()).DoAndReturn(
		func(id int) (*User, error) {
			if id < 0 {
				return nil, errors.New("invalid ID")
			}
			return &User{ID: id, Name: "User " + string(rune(id))}, nil
		},
	).AnyTimes()

	service := NewUserService(mockRepo)

	// Test with valid ID
	name, err := service.GetUserName(1)
	if err != nil {
		t.Errorf("Expected no error for valid ID, got %v", err)
	}
	if name == "" {
		t.Error("Expected non-empty name")
	}

	// Test with invalid ID
	_, err = service.GetUserName(-1)
	if err == nil {
		t.Error("Expected error for invalid ID, got nil")
	}
}

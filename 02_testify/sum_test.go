package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestSum demonstrates basic assertions with testify/assert.
func TestSum(t *testing.T) {
	result := Sum(2, 3)

	// assert.Equal continues execution even if it fails
	assert.Equal(t, 5, result, "Sum(2, 3) should equal 5")
	assert.NotEqual(t, 6, result, "Sum(2, 3) should not equal 6")
}

// TestSum_WithRequire demonstrates require (fails fast).
func TestSum_WithRequire(t *testing.T) {
	result := Sum(2, 3)

	// require.Equal stops execution immediately if it fails
	require.Equal(t, 5, result, "Sum must equal 5")

	// This line won't execute if the require above fails
	assert.Greater(t, result, 0, "Result should be positive")
}

// TestSum_TableDriven demonstrates table-driven tests with testify.
func TestSum_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -2, -3},
		{"with zero", 0, 5, 5},
		{"both zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestMultiply demonstrates various assertion types.
func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)

	assert.Equal(t, 12, result)
	assert.NotZero(t, result)
	assert.Greater(t, result, 10)
	assert.Less(t, result, 15)
	assert.IsType(t, 0, result) // verify it's an int
}

// TestDivide demonstrates error assertions.
func TestDivide(t *testing.T) {
	t.Run("successful division", func(t *testing.T) {
		result, err := Divide(6, 2)

		assert.NoError(t, err)
		assert.Equal(t, 3, result)
	})

	t.Run("division by zero returns error", func(t *testing.T) {
		result, err := Divide(5, 0)

		assert.Error(t, err)
		assert.Equal(t, ErrDivisionByZero, err)
		assert.Zero(t, result)
	})
}

// TestDivide_WithRequire demonstrates require for critical checks.
func TestDivide_WithRequire(t *testing.T) {
	result, err := Divide(10, 2)

	// If this fails, stop immediately - no point continuing
	require.NoError(t, err, "Division should not produce an error")

	// Only reaches here if no error occurred
	assert.Equal(t, 5, result)
}

// TestCollectionAssertions demonstrates assertions for slices and maps.
func TestCollectionAssertions(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	assert.Len(t, numbers, 5)
	assert.Contains(t, numbers, 3)
	assert.NotContains(t, numbers, 10)
	assert.ElementsMatch(t, []int{5, 4, 3, 2, 1}, numbers) // order doesn't matter

	userMap := map[string]int{"alice": 25, "bob": 30}
	assert.Contains(t, userMap, "alice")
	assert.Equal(t, 25, userMap["alice"])
}

// TestNilAssertions demonstrates nil checking.
func TestNilAssertions(t *testing.T) {
	var ptr *int
	assert.Nil(t, ptr)

	value := 42
	ptr = &value
	assert.NotNil(t, ptr)
	assert.Equal(t, 42, *ptr)
}

// MockUserRepository is a mock implementation of UserRepository.
type MockUserRepository struct {
	mock.Mock
}

// GetUser is the mocked method.
func (m *MockUserRepository) GetUser(id int) (*User, error) {
	args := m.Called(id)

	// Handle nil return value
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*User), args.Error(1)
}

// SaveUser is the mocked method.
func (m *MockUserRepository) SaveUser(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

// TestUserService_WithMock demonstrates mocking with testify/mock.
func TestUserService_WithMock(t *testing.T) {
	t.Run("get user name successfully", func(t *testing.T) {
		// Create a new mock
		mockRepo := new(MockUserRepository)

		// Set expectations
		expectedUser := &User{ID: 1, Name: "John Doe"}
		mockRepo.On("GetUser", 1).Return(expectedUser, nil)

		// Create service with mock
		service := NewUserService(mockRepo)

		// Call the method being tested
		name, err := service.GetUserName(1)

		// Assert results
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", name)

		// Verify that all expectations were met
		mockRepo.AssertExpectations(t)
	})

	t.Run("get user returns error", func(t *testing.T) {
		mockRepo := new(MockUserRepository)

		// Simulate an error
		mockRepo.On("GetUser", 999).Return(nil, Error("user not found"))

		service := NewUserService(mockRepo)
		name, err := service.GetUserName(999)

		assert.Error(t, err)
		assert.Empty(t, name)
		assert.Contains(t, err.Error(), "not found")

		mockRepo.AssertExpectations(t)
	})
}

// TestMock_AdvancedUsage demonstrates advanced mock features.
func TestMock_AdvancedUsage(t *testing.T) {
	mockRepo := new(MockUserRepository)

	// Mock can match any argument
	mockRepo.On("GetUser", mock.Anything).Return(&User{ID: 1, Name: "Any User"}, nil)

	service := NewUserService(mockRepo)

	// Works with any ID
	name1, _ := service.GetUserName(1)
	name2, _ := service.GetUserName(100)

	assert.Equal(t, "Any User", name1)
	assert.Equal(t, "Any User", name2)

	// Verify GetUser was called exactly twice
	mockRepo.AssertNumberOfCalls(t, "GetUser", 2)
}

// TestMock_VerifyCallOrder demonstrates call order verification.
func TestMock_VerifyCallOrder(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &User{ID: 1, Name: "John"}
	mockRepo.On("GetUser", 1).Return(user, nil).Once()
	mockRepo.On("SaveUser", user).Return(nil).Once()

	// Call in specific order
	_, _ = mockRepo.GetUser(1)  // Error ignored - testing mock behavior
	_ = mockRepo.SaveUser(user) // Error ignored - testing mock behavior

	// Verify calls were made
	mockRepo.AssertCalled(t, "GetUser", 1)
	mockRepo.AssertCalled(t, "SaveUser", user)
	mockRepo.AssertExpectations(t)
}

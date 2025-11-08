package testify

// Sum returns the sum of two integers.
func Sum(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the result of dividing a by b.
// Returns an error if b is zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// ErrDivisionByZero is returned when attempting to divide by zero.
var ErrDivisionByZero = Error("division by zero")

// Error represents a simple error type.
type Error string

func (e Error) Error() string {
	return string(e)
}

// User represents a user in the system.
type User struct {
	ID   int
	Name string
}

// UserRepository defines the interface for user data access.
type UserRepository interface {
	GetUser(id int) (*User, error)
	SaveUser(user *User) error
}

// UserService provides business logic for user operations.
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUserName retrieves a user's name by ID.
func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

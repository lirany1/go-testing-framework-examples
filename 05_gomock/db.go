package gomock

//go:generate mockgen -source=db.go -destination=mock_db.go -package=gomock

// User represents a user in the system.
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	// GetUser retrieves a user by ID.
	GetUser(id int) (*User, error)

	// SaveUser persists a user to the database.
	SaveUser(user *User) error

	// DeleteUser removes a user by ID.
	DeleteUser(id int) error

	// ListUsers returns all users.
	ListUsers() ([]*User, error)
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

// CreateUser creates a new user.
func (s *UserService) CreateUser(name, email string) error {
	user := &User{
		Name:  name,
		Email: email,
	}
	return s.repo.SaveUser(user)
}

// RemoveUser deletes a user by ID.
func (s *UserService) RemoveUser(id int) error {
	return s.repo.DeleteUser(id)
}

// GetAllUserNames returns the names of all users.
func (s *UserService) GetAllUserNames() ([]string, error) {
	users, err := s.repo.ListUsers()
	if err != nil {
		return nil, err
	}

	names := make([]string, len(users))
	for i, user := range users {
		names[i] = user.Name
	}
	return names, nil
}

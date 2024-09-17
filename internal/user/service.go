package user

import "github.com/cemtanrikut/user-service/cmd/user-service/main.go/pkg/utils"

// UserService handles user operations
type UserService struct {
	repo *UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(firstName, lastName, nickname, email, password, country string) (User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return User{}, err
	}

	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Password:  hashedPassword,
		Email:     email,
		Country:   country,
	}

	return s.repo.AddUser(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id, firstName, lastName, nickname, email, country string) (User, error) {
	// Kullanıcıyı repository'den çekiyoruz
	existingUser, err := s.repo.GetUser(id)
	if err != nil {
		return User{}, err
	}

	// Kullanıcı bilgilerini güncelliyoruz
	updatedUser := User{
		ID:        existingUser.ID,
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Password:  existingUser.Password, // don't change password
		Email:     email,
		Country:   country,
		CreatedAt: existingUser.CreatedAt,
	}

	return s.repo.UpdateUser(id, updatedUser)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

// ListUsers returns a list of users with pagination and filtering by country
func (s *UserService) ListUsers(country string, limit, offset int) []User {
	return s.repo.ListUsers(country, limit, offset)
}

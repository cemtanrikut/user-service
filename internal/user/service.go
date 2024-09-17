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

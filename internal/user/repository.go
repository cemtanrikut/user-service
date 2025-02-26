package user

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

// UserRepository provides in-memory storage for users
type UserRepository struct {
	mu    sync.Mutex
	users map[string]User
}

// NewUserRepository creates a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]User),
	}
}

// AddUser adds a new user to the repository
func (r *UserRepository) AddUser(u User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// create new UUID
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	r.users[u.ID] = u
	return u, nil
}

// GetUser retrieves a user by ID
func (r *UserRepository) GetUser(id string) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser updates an existing user in the repository
func (r *UserRepository) UpdateUser(id string, updatedUser User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	// update process
	updatedUser.ID = user.ID
	updatedUser.CreatedAt = user.CreatedAt
	updatedUser.UpdatedAt = time.Now()

	r.users[id] = updatedUser
	return updatedUser, nil
}

// DeleteUser deletes a user by ID
func (r *UserRepository) DeleteUser(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}

// ListUsers returns a paginated list of users with optional filtering by any field
func (r *UserRepository) ListUsers(filters map[string]string, limit, offset int) []User {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []User
	for _, user := range r.users {
		match := true

		// all filters
		for key, value := range filters {
			switch key {
			case "first_name":
				if user.FirstName != value {
					match = false
				}
			case "last_name":
				if user.LastName != value {
					match = false
				}
			case "nickname":
				if user.Nickname != value {
					match = false
				}
			case "email":
				if user.Email != value {
					match = false
				}
			case "country":
				if user.Country != value {
					match = false
				}
			}
		}

		if match {
			result = append(result, user)
		}
	}

	// pagination
	start := offset
	end := offset + limit

	if start > len(result) {
		return []User{}
	}

	if end > len(result) {
		end = len(result)
	}

	return result[start:end]
}

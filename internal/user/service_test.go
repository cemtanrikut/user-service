package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test for creating a new user
func TestCreateUser(t *testing.T) {
	repo := NewUserRepository()
	service := NewUserService(repo)

	user, err := service.CreateUser("Alice", "Bob", "AB123", "alice@bob.com", "password123", "UK")

	assert.NoError(t, err)
	assert.NotNil(t, user.ID)
	assert.Equal(t, "Alice", user.FirstName)
	assert.Equal(t, "Bob", user.LastName)
	assert.Equal(t, "AB123", user.Nickname)
	assert.Equal(t, "alice@bob.com", user.Email)
	assert.Equal(t, "UK", user.Country)
	assert.WithinDuration(t, time.Now(), user.CreatedAt, time.Second)
}

// Test for updating an existing user
func TestUpdateUser(t *testing.T) {
	repo := NewUserRepository()
	service := NewUserService(repo)

	// create first user
	user, _ := service.CreateUser("Alice", "Bob", "AB123", "alice@bob.com", "password123", "UK")

	// update user
	updatedUser, err := service.UpdateUser(user.ID, "AliceUpdated", "BobUpdated", "AB123Updated", "alice.updated@bob.com", "US")

	assert.NoError(t, err)
	assert.Equal(t, "AliceUpdated", updatedUser.FirstName)
	assert.Equal(t, "BobUpdated", updatedUser.LastName)
	assert.Equal(t, "AB123Updated", updatedUser.Nickname)
	assert.Equal(t, "alice.updated@bob.com", updatedUser.Email)
	assert.Equal(t, "US", updatedUser.Country)
	assert.Equal(t, user.ID, updatedUser.ID) // ID aynı olmalı
}

// Test for deleting a user
func TestDeleteUser(t *testing.T) {
	repo := NewUserRepository()
	service := NewUserService(repo)

	// create new user
	user, _ := service.CreateUser("Alice", "Bob", "AB123", "alice@bob.com", "password123", "UK")

	// delete user
	err := service.DeleteUser(user.ID)

	assert.NoError(t, err)

	// check user not found
	_, err = service.repo.GetUser(user.ID)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}

// Test for listing users with pagination and filtering
func TestListUsers(t *testing.T) {
	repo := NewUserRepository()
	service := NewUserService(repo)

	// create several user
	_, _ = service.CreateUser("Alice", "Bob", "AB123", "alice@bob.com", "password123", "UK")
	_, _ = service.CreateUser("John", "Doe", "JD456", "john@doe.com", "password456", "US")
	_, _ = service.CreateUser("Jane", "Doe", "JD789", "jane@doe.com", "password789", "UK")

	// list users from country = UK
	filters := map[string]string{"country": "UK"}
	users := service.ListUsers(filters, 10, 0)

	assert.Len(t, users, 2) // should be 2 user from UK
	assert.Equal(t, "UK", users[0].Country)
	assert.Equal(t, "UK", users[1].Country)
}

// Test for filtering users by first name
func TestListUsersByFirstName(t *testing.T) {
	repo := NewUserRepository()
	service := NewUserService(repo)

	// create several user
	_, _ = service.CreateUser("Alice", "Bob", "AB123", "alice@bob.com", "password123", "UK")
	_, _ = service.CreateUser("Alice", "Smith", "AS123", "alice@smith.com", "password456", "US")

	// list Alice users
	filters := map[string]string{"first_name": "Alice"}
	users := service.ListUsers(filters, 10, 0)

	assert.Len(t, users, 2) // should be 2 Alice named user
	assert.Equal(t, "Alice", users[0].FirstName)
	assert.Equal(t, "Alice", users[1].FirstName)
}

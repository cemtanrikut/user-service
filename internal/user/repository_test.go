package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for adding and retrieving a user
func TestAddAndGetUser(t *testing.T) {
	repo := NewUserRepository()

	user, err := repo.AddUser(User{
		FirstName: "Alice",
		LastName:  "Bob",
		Nickname:  "AB123",
		Email:     "alice@bob.com",
		Country:   "UK",
	})

	assert.NoError(t, err)
	assert.NotNil(t, user.ID)

	retrievedUser, err := repo.GetUser(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, retrievedUser.ID)
	assert.Equal(t, "Alice", retrievedUser.FirstName)
	assert.Equal(t, "Bob", retrievedUser.LastName)
}

// Test for deleting a user
func Test_DeleteUser(t *testing.T) {
	repo := NewUserRepository()

	user, _ := repo.AddUser(User{
		FirstName: "Alice",
		LastName:  "Bob",
		Nickname:  "AB123",
		Email:     "alice@bob.com",
		Country:   "UK",
	})

	err := repo.DeleteUser(user.ID)
	assert.NoError(t, err)

	_, err = repo.GetUser(user.ID)
	assert.Error(t, err)
}

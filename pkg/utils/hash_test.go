package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for hashing a password
func TestHashPassword(t *testing.T) {
	plainPassword := "supersecurepassword"

	hashedPassword, err := HashPassword(plainPassword)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	assert.NotEqual(t, plainPassword, hashedPassword)
}

// Test for checking a valid password
func TestCheckPasswordValid(t *testing.T) {
	plainPassword := "supersecurepassword"
	hashedPassword, _ := HashPassword(plainPassword)

	err := CheckPassword(plainPassword, hashedPassword)
	assert.NoError(t, err)
}

func TestCheckPasswordInvalid(t *testing.T) {
	plainPassword := "supersecurepassword"
	hashedPassword, _ := HashPassword(plainPassword)

	err := CheckPassword("wrongpassword", hashedPassword)
	assert.Error(t, err)
}

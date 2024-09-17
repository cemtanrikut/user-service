package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Test for creating a new user (HTTP POST /users)
func TestCreateUserHandler(t *testing.T) {
	repo := NewUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	// Test HTTP request
	payload := `{
		"first_name": "Alice",
		"last_name": "Bob",
		"nickname": "AB123",
		"email": "alice@bob.com",
		"password": "password123",
		"country": "UK"
	}`

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users", handler.CreateUserHandler).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var user User
	json.NewDecoder(rr.Body).Decode(&user)
	assert.Equal(t, "Alice", user.FirstName)
	assert.Equal(t, "Bob", user.LastName)
	assert.Equal(t, "AB123", user.Nickname)
	assert.Equal(t, "UK", user.Country)
}

// Test for listing users (HTTP GET /users)
func TestListUsersHandler(t *testing.T) {
	repo := NewUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	// Some mock users
	_, _ = service.CreateUser("Alice", "Bob", "AB123", "alice@bob.com", "password123", "UK")
	_, _ = service.CreateUser("John", "Doe", "JD456", "john@doe.com", "password456", "US")

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users", handler.ListUsersHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var users []User
	json.NewDecoder(rr.Body).Decode(&users)

	assert.Len(t, users, 2)
	assert.Equal(t, "Alice", users[0].FirstName)
	assert.Equal(t, "John", users[1].FirstName)
}

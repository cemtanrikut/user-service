package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// UserHandler provides HTTP handlers for user operations
type UserHandler struct {
	service *UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUserHandler handles creating a new user
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Nickname  string `json:"nickname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Country   string `json:"country"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(req.FirstName, req.LastName, req.Nickname, req.Email, req.Password, req.Country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler handles updating an existing user
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Nickname  string `json:"nickname"`
		Email     string `json:"email"`
		Country   string `json:"country"`
	}

	vars := mux.Vars(r)
	userID := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.service.UpdateUser(userID, req.FirstName, req.LastName, req.Nickname, req.Email, req.Country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUserHandler handles deleting a user
func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := h.service.DeleteUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

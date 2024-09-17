package main

import (
	"log"
	"net/http"

	"github.com/cemtanrikut/user-service/cmd/user-service/main.go/internal/user"
	"github.com/gorilla/mux"
)

func main() {
	repo := user.NewUserRepository()
	service := user.NewUserService(repo)
	handler := user.NewUserHandler(service)

	// generate route
	r := mux.NewRouter()

	// Endpoint for health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Create user
	r.HandleFunc("/users", handler.CreateUserHandler).Methods("POST")

	// Update user
	r.HandleFunc("/users/{id}", handler.UpdateUserHandler).Methods("PUT")

	// Delete user
	r.HandleFunc("/users/{id}", handler.DeleteUserHandler).Methods("DELETE")

	// Filter and list users
	r.HandleFunc("/users", handler.ListUsersHandler).Methods("GET")

	// Start server
	log.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

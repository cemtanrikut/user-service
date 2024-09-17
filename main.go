package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Router oluşturuyoruz
	r := mux.NewRouter()

	// Endpoint for health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start server
	log.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

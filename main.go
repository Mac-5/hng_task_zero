package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Email       string `json:"email"`
	CurrentTime string `json:"datetime"`
	GitHubURL   string `json:"github_repo"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Email:       "emmamarkpius@gmail.com",
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
		GitHubURL:   "https://github.com/Mac-5/hng_task_zero",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)    // Chi's built-in logger middleware
	r.Use(middleware.Recoverer) // Graceful panic recovery

	// Routes
	r.Get("/api/", apiHandler)

	// Start server
	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

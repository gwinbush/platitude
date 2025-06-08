package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
)

type (
	PlatitudeResponse struct {
		Message string `json:"message"`
	}

	HealthResponse struct {
		Status string `json:"status"`
	}
)

func platitudeHandler(w http.ResponseWriter, r *http.Request) {
	var platitudes = []string{
		"Keep it simple, stupid.",
		"Don't communicate by sharing memory, share memory by communicating.",
		"Make it work, make it right, make it fast.",
		"Less is more.",
		"Clear is better than clever.",
	}

	randomIndex := rand.IntN(len(platitudes))
	response := PlatitudeResponse{
		Message: platitudes[randomIndex],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{Status: "healthy"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/golang/platitude", platitudeHandler)
	http.HandleFunc("/health", healtHandler)

	port := os.Getenv("PORT")
	fmt.Println("Server is listening on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

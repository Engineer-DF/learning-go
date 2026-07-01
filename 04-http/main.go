package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done"`
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&tasks); err != nil {
		fmt.Printf("failed to encode JSON: %v\n", err)
	}
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	var newTask task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	// add functionality later
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// add functionality later
}

var tasks = []task{
	{ID: 0, Title: "Write simple HTTP API", Description: "This is quite difficult for me", Done: false},
	{ID: 1, Title: "Check my API", Description: "I hope it works", Done: false},
	{ID: 2, Title: "Fix any bugs", Done: false},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", GetTask)
	mux.HandleFunc("POST /tasks", PostTask)

	http.ListenAndServe(":8080", mux)
}

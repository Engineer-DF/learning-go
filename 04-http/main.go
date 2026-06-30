package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(users); err != nil {
			fmt.Printf("failed to encode JSON: %v\n", err)
		}

	case http.MethodPost:
		var newUser User

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "bad request body", http.StatusBadRequest)
			return
		}

		newUser.ID = len(users)
		users = append(users, newUser)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newUser); err != nil {
			fmt.Printf("failed to encode JSON %V\n", err)
		}

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

	}
}

var users = []User{
	{
		ID:   0,
		Name: "Serega Magnum",
		Age:  19,
	},
	{
		ID:   1,
		Name: "Alice Fox",
		Age:  20,
	},
	{
		ID:   2,
		Name: "Alpamys Kazakhski",
		Age:  34,
	},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}

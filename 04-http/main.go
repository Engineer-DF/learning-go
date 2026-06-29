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
	if r.Method == http.MethodGet {
		_, err := http.Get("")
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
		//w.Write([]byte(`{"status":"ok"}`))
	} else {
		fmt.Println("Unexpected action!")
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

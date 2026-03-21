package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var users []User

	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	fmt.Println("Received Successfully:")
	for i, v := range users {
		fmt.Printf("%d -> Name: %s, Age: %d\n", i, v.Name, v.Age)
	}

	w.Write([]byte("Users Received Successfully..."))
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:9080",
	}

	fmt.Println("Listen on :9080")

	http.HandleFunc("/js", jsonHandler)

	server.ListenAndServe()
}

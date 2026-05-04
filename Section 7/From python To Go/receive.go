package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Payload struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	IsActive bool     `json:"is_active"`
	Score    float64  `json:"score"`
	Skills   []string `json:"skills"`
	Message  string   `json:"message"`
}

func PyHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var pyjson Payload
	// err := json.NewDecoder(r.Body).Decode(&pyjson)
	json.Unmarshal(body, &pyjson)
	// if err != nil {
	// 	http.Error(w, "Invalid Json", http.StatusBadRequest)
	// 	return
	// }

	fmt.Printf("Name: %s, Age: %d, Active: %v\n", pyjson.Name, pyjson.Age, pyjson.IsActive)
	fmt.Printf("Score: %.1f, Skills: %v\n", pyjson.Score, pyjson.Skills)
	fmt.Printf("Message: %s\n", pyjson.Message)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok", "message": "Data received by Go!"}`))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}
	fmt.Println("Listen on: 8080")
	http.HandleFunc("/receive", PyHandler)
	server.ListenAndServe()
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error in Reading Body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	msg := string(body)

	fmt.Println("The message received is: ", msg)

	w.Write([]byte("Message received successfully"))

}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File is too large", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Couldn't retrive file", http.StatusBadRequest)
		return
	}

	file_name := handler.Filename

	dst, _ := os.Create(file_name)

	defer dst.Close()

	io.Copy(dst, file)

	fmt.Println("File uploaded successfully")

	w.Write([]byte("File uploaded successfully"))
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Allowed", http.StatusMethodNotAllowed)
		return
	}

	var users []User

	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	for i, u := range users {
		fmt.Printf("The data of the %dth user is: %v\n", i, u)
	}

}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:9080",
	}

	http.HandleFunc("/rec", JsonHandler)

	fmt.Println("Listen on: 9080")

	server.ListenAndServe()
}

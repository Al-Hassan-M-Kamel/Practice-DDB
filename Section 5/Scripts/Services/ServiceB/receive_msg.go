package main

import (
	"fmt"
	"io"
	"net/http"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error in reading message", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	msg := string(body)
	fmt.Println("Received Message: ", msg)

	response := "Message Received"
	w.Write([]byte(response))

}

func main2() {

	server := http.Server{
		Addr: "127.0.0.1:9080",
	}

	fmt.Println("Listen on:9080")
	http.HandleFunc("/msg", MessageHandler)

	server.ListenAndServe()
}

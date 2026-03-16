package main

import (
	"fmt"
	"io"
	"net/http"
)

func CallServiceA(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Fprintf(w, "Service B received: %s", string(body))
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:9080",
	}

	http.HandleFunc("/callA", CallServiceA)

	fmt.Println("Service B running on :9080")

	server.ListenAndServe()
}

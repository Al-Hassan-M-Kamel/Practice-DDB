package main

import (
	"fmt"
	"io"
	"net/http"
)

func Handler1(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World...")
}

func Handler2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World from handler2...")
}

func main() {

	resp, err := http.Get("http://127.0.0.1:9080/hi")

	if err != nil{
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("The reponser form Service is: %s\n", string(body))



}

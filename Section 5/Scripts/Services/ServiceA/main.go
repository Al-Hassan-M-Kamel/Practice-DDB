package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From ServiceA")
}

func main1() {

	server := http.Server{
		Addr: "127.0.0.1:1010",
	}

	fmt.Println("ServiceA running on :1010")
	http.HandleFunc("/hello", HelloHandler)

	server.ListenAndServe()

}

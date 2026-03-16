package main

import (
	"fmt"
	"net/http"
	"time"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello..!")
	fmt.Println("Start new request from hello...")
	for i := 0; i < 10; i++ {
		fmt.Println("Hello..!")
		time.Sleep(1.5 * 1e9)
	}
	fmt.Println("End of the request from hello...")

}

type WorldHandler struct{}

func (wo *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World..!")
	fmt.Println("Start new request from world...")
	for i := 0; i < 10; i++ {
		fmt.Println("World..!")
		// time.Sleep(1.5 * 1e9)
	}
	fmt.Println("End of the request from world...")
}
func main() {

	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:9080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}

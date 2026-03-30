package main

import (
	"Scripts/Services/SServices"
	"fmt"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: "127.0.0.1:9080",
	}

	http.HandleFunc("/save", SServices.Save_Handler)

	fmt.Println("Listen on: 9080")

	server.ListenAndServe()

}

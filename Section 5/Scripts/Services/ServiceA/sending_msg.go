package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main2() {

	msg := fmt.Sprintf("Hello From Client number: %d", 5)

	resp, err := http.Post("http://127.0.0.1:9080/msg", "text/plain", bytes.NewBuffer([]byte(msg)))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Sent Successfully...")
}

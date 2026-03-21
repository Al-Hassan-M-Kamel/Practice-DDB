package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	users := []User{
		{Name: "Ali", Age: 25},
		{Name: "Omar", Age: 30},
		{Name: "Sara", Age: 22},
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://127.0.0.1:9080/js", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Sent Successfully...")

}

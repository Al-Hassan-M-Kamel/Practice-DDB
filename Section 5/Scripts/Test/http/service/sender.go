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

	json_data, _ := json.Marshal(users)

	resp, err := http.Post("http://127.0.0.1:9080/rec", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Done Sending...")
}

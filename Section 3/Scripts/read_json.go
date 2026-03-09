package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address2 struct {
	Type    string `json:"type"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type VCard2 struct {
	FirstName string      `json:"firstname"`
	LastName  string      `json:"lastname"`
	Addresses []*Address2 `json:"addresses"`
	Remark    string      `json:"remark"`
}

func main3() {

	// First read the content of the file as bytes...
	// This is suitable with for small/medium json files...
	data, err := os.ReadFile("vcard.json")

	if err != nil {
		panic(err)
	}

	var vc VCard2

	err = json.Unmarshal(data, &vc)
	if err != nil {
		panic(err)
	}

	fmt.Println(vc)

}

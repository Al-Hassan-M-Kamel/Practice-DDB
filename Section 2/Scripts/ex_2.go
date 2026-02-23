package main

import (
	"fmt"
	"io"
)

func Ex_2() {

	var slice []int
	var number int

	for {

		fmt.Print("Please Enter a number or ^Z to end input: ")
		_, err := fmt.Scan(&number)
		if err == io.EOF { // check that the user has input ^Z or not...
			break
		} else if err != nil { // if err == nil then the input is correct and matches the variable...
			fmt.Println("You have entered non matching input...")
			break
		}
		slice = append(slice, number)
		fmt.Printf("At this iteration slice length is: %d and its capacity is: %d\n", len(slice), cap(slice))
	}

}

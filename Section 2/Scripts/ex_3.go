package main

import (
	"fmt"
	"io"
	"sort"
)

func Remove_Duplicates(slice []int) []int {

	i := 1
	for i < len(slice) {
		if slice[i] == slice[i-1] {
			slice = append(slice[0:i], slice[i+1:]...)
			i = 1
			continue
		}
		i++
	}
	return slice

}

func EX_3() {

	var slice []int
	var number int

	// Take input from user...
	for {
		fmt.Print("Enter a number or ^Z to end the input: ")
		_, err := fmt.Scan(&number)
		if err == io.EOF {
			break
		}

		slice = append(slice, number)
	}

	// Sort the slice...
	sort.Ints(slice)

	fmt.Printf("The Sorted Slice: %v", slice)

	slice = Remove_Duplicates(slice)

	fmt.Printf("The slice without duplicates: %v", slice)
}

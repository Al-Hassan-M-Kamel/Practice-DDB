/*

Problem: Implement an in‑place array rotation by k positions to the right (e.g., [1,2,3,4,5] rotated by 2 becomes [4,5,1,2,3]).
Handle k > len(arr) and k < 0.

Input: Array of 10 integers, k from stdin.
Output: Rotated array, space‑separated.
Example:

text
Input: 1 2 3 4 5 3 ; k=1
Output: 3 1 2 3 4 5

*/

package main

import (
	"fmt"
)

func main() {

	// Take input from user and declare variables...
	var array_size, k int

	fmt.Println("Please enter array size:")
	fmt.Scan(&array_size)

	fmt.Println("Please enter k:")
	fmt.Scan(&k)

	// Take the elements of the array...
	arr := make([]int, array_size)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Enter the %dth element:", i)
		fmt.Scan(&arr[i])
		fmt.Println()
	}

	// Handle k:
	k %= array_size // k > array_size
	if k < 0 {
		k += array_size
	}

	result := make([]int, array_size) // the new array that will hold the rotated values...
	index := array_size - k           // the offset that will be used to address the original array...

	for i, _ := range arr {
		result[i] = arr[index]
		index = (index + 1) % array_size // this updates makes sure that the index will return to zero when reach the limit...
	}

	fmt.Println(arr)

	/*
			*** CHALLENGE YOURSELF ***
			this code has a memory complexit of O(2) because we used two array try to rewrite the code so that it will be O(1)
			// Rotation number one from 0 to array_size-1
		end := array_size -1
		for i :=0; i < end; i++{
			arr[i], arr[end] = arr[end], arr[i]
			end --
		}

		// Rotation number two from 0 to k-1
		end = k-1
		for i :=0; i < end; i++{
			arr[i], arr[end] = arr[end], arr[i]
			end --
		}

		// Rotation number three from k to array_size - 1
		end = array_size - 1
		for i :=k; i < end; i++{
			arr[i], arr[end] = arr[end], arr[i]
			end --
		}


	*/

}

package main

import "fmt"

func Ex3BranchlessBitwise() {

	// Take inputs from user...
	fmt.Println("Enter the array size: ")
	var array_size int
	fmt.Scan(&array_size)

	arr := make([]int, array_size)

	for i := 0; i < array_size; i++ {
		fmt.Printf("Enter the %dth element: ", i)
		fmt.Scan(&arr[i])

	}

	// Set tracker variables...
	max_sum := arr[0]
	current_sum := arr[0]
	start, end, temp_start := 0, 0, 0

	for i := 1; i < array_size; i++ {

		// check whether the current sum greater that the current value ...
		checkCurrent := ((-current_sum) >> 63)
		// no true and current sum is smaller than the current value then make a fresh start...
		current_sum = (arr[i] &^ checkCurrent) | ((current_sum + arr[i]) & checkCurrent) 
		temp_start = (i &^ checkCurrent) | (temp_start & checkCurrent)
		// update max sum with new maximum value...
		checkMax := ((current_sum - max_sum) >> 63)
		max_sum = (current_sum &^ checkMax) | (max_sum & checkMax)
		start, end = (temp_start &^ checkMax) | (start & checkMax), (i &^ checkMax) | (end & checkMax)
	
	}

	fmt.Printf("Max sum: %d from index %d to %d\n", max_sum, start, end)

}

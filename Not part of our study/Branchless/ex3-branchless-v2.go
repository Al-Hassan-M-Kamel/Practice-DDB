package main

import "fmt"

func Ex3BranchlessMap() {

	// Take inputs from user...
	fmt.Println("Enter the array size: ")
	var array_size int
	fmt.Scan(&array_size)

	arr := make([]int, array_size)
	for i := 0; i < array_size; i++ {
		fmt.Printf("Enter the %dth element: ", i)
		fmt.Scan(&arr[i])

	}
	// adding a map to replace the if-else statements
	check := map[bool]int{
		true:  1,
		false: 0,
	}

	// Set tracker variables...
	max_sum := arr[0]
	current_sum := arr[0]
	start, end, temp_start := 0, 0, 0

	for i := 1; i < array_size; i++ {

		old_current_sum := current_sum
		// check whether the current sum greater that the current value ...
		// no true and current sum is smaller than the current value then make a fresh start...
		current_sum = (arr[i]*check[old_current_sum<0]) + ((current_sum+arr[i])*check[old_current_sum>=0])
		temp_start = (i*check[old_current_sum<0]) + (temp_start*check[old_current_sum>=0])

		old_max_sum := max_sum
		// update max sum with new maximum value...
		max_sum = (current_sum*check[current_sum>old_max_sum]) + (old_max_sum*check[current_sum<=old_max_sum])
		start, end = (temp_start*check[current_sum>old_max_sum]) + (start*check[current_sum<=old_max_sum]), (i*check[current_sum>old_max_sum]) + (end*check[current_sum<=old_max_sum])
	}

	fmt.Printf("Max sum: %d from index %d to %d\n", max_sum, start, end)

}

package main

import "fmt"

func Sum_of_Digits(number int) int {

	if number < 10 {
		return number
	}

	return number%10 + Sum_of_Digits(number/10)

}

func Index_Of(arr []int, target int, index int) int {

	if len(arr) == index {
		return -1
	}

	if arr[index] == target {
		return index
	}

	return Index_Of(arr, target, index+1)

}

func Fibonacci(n ...int) {

	if len(n) != 3 {
		n = append(n, 0, 1)
	}

	if n[0] < 0 {
		return
	}
	fmt.Printf("%d ", n[1])
	Fibonacci(n[0]-1, n[2], n[1]+n[2])
}

func EX_5() {

	// arr := []int{1, 3, 3, 2, 5}

	// fmt.Println(Index_Of(arr, 1, 2))

	Fibonacci(9)
}

package main

import "fmt"

func hello() {

	var input_value int
	max := 0
	for i := 0; i < 10; i++ {
		fmt.Scan(&input_value)
		if input_value > max{
			max = input_value
		}
	}	

	fmt.Println(max)

}

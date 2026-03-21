package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {

	fmt.Println("In main()")

	ch := make(chan string, 2)

	go Long_Wait(ch)

	go Short_Wait(ch)

	// fmt.Println("About to sleep in main()")
	// time.Sleep(10 * 1e9)

	fmt.Println("main() is waiting...")

	// Waiting the goroutines to finish...
	for i := 0; i < 2; i++ {
		<-ch
	}

	fmt.Println("At the end of the main()")

}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func Send_Data(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func Get_Data(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func Long_Wait(ch chan string) {
	fmt.Println("Beginning Long_Wait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of Long_Wait()")

	ch <- "I have finished"
}

func Short_Wait(ch chan string) {
	fmt.Println("Beginning of Short_Wait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of Short_Wait()")

	ch <- "I have finished"
}


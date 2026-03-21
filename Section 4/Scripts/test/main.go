package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("In main()")

	ch := make(chan string, 2)

	go Long_Wait(ch)

	go Short_Wait(ch)

	for i := 0; i < 2; i++ {
		<-ch
	}

	fmt.Println("main() has finished")

}
func pump(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func Send_Data(ch chan string) {
	ch <- "Asmaa"
	ch <- "Wakeel"
	ch <- "Ismail"
	ch <- "Ahmed"
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
	time.Sleep(5 * 1e9)
	fmt.Println("End of Long_Wait()")

	ch <- "I have finished"
}

func Short_Wait(ch chan string) {
	fmt.Println("Beginning Short_Wait()")
	time.Sleep(2 * 1e9)
	fmt.Println("End of Short_Wait()")

	ch <- "I have finished"
}

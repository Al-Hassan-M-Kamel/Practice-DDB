package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ex_4() {

	dict := make(map[string][]int)

	// first define scanner which will be corresponding to a specific stream...
	scanner := bufio.NewScanner(os.Stdin)

	// Now take input from user until the new line works like (Console.ReadLine() in C#)...
	scanner.Scan()

	sentence := scanner.Text() // It's the actual input...
	// if the input is single number you can use strConv.ParseInt(input)...

	words := strings.Split(sentence, " ")

	for i, word := range words {
		dict[word] = append(dict[word], i)
	}

	fmt.Println(dict)

}

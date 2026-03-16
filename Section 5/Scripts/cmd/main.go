package main

import (
	"bufio"
	"flag" //  implements command-line flag parsing.
	"fmt"
	"io"
	"os"
	"strings"
)

func Test() {
	// Define command line arguments...
	// Argument number 1:
	name := flag.String("name", "world", "the name to greet")
	// This declares a string flag, -name, stored in the pointer name, with type *string

	// Argument number 2:
	upper := flag.Bool("upper", false, "print in upper case")
	// This declares a boolean flag, -upper, stored in the pointer upper, with type *bool

	// After all flags are defined call:
	flag.Parse() // to parse the command line into the defined flags.

	msg := fmt.Sprintf("Hello, %s", *name)
	if *upper {
		msg = strings.ToUpper(msg)
	}

}

func main() {

	in := flag.String("in", "", "Take the path to the fasta file")

	flag.Parse()

	fmt.Println(Count_Nucleotieds(*in))

}

func Count_Nucleotieds(input_file string) map[string]int {
	/*
		This function takes as input a fasta file and it counts the number of nucleotide occurrences in this file
		and it returns a map for each nucleotide count...

		input_file: takes the fasta file path...
	*/

	in_file, err := os.Open(input_file)

	if err != nil {
		fmt.Println("Error in opening the file")
		return nil
	}

	defer in_file.Close()

	reader := bufio.NewReader(in_file)

	word_count := make(map[string]int)

	var line string
	reader.ReadString('\n') // read and skip the first line
	for {
		line, err = reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		for _, nt := range strings.TrimSpace(line) {

			word_count[string(nt)]++
		}
	}

	return word_count
}

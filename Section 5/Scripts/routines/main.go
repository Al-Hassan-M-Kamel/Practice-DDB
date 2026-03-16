package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// Where the files are located...
	data_dir := "F:\\Work\\Courses\\DDB\\Material\\Section 5\\Scripts\\Data"

	var file_name string

	// Read the content of the folder, files will contain all the file names in the directory...
	files, err := os.ReadDir(data_dir)
	if err != nil {
		panic(err)
	}

	// This channel is responsilbe for data flow between routines...
	flow := make(chan map[string]int, len(files))

	// This channel is responsilbe for index flow between routines...
	index_flow := make(chan int, len(files))

	// Create the output file...
	out_file, err := os.Create("./routines/output.txt")
	if err != nil {
		panic(err)
	}

	defer out_file.Close()

	writer := bufio.NewWriter(out_file)

	// Creating stacks for each file...
	for index, file := range files {

		file_name = data_dir + "//" + file.Name()
		go Count_Nucleotieds(file_name, flow, index_flow, index)
	}

	var mp map[string]int
	var ind int

	// Recieving results from the routines...
	for i := 0; i < len(files); i++ {

		mp = <-flow
		ind = <-index_flow

		// Writing data into the output file...
		fmt.Fprintf(writer, "Results of file with index: %d \n", ind)
		for k, v := range mp {
			fmt.Fprintf(writer, "%s: %d\n", k, v)
		}

		fmt.Fprintln(writer)

	}

	writer.Flush()

}

func Count_Nucleotieds(input_file string, flow chan map[string]int, index_flow chan int, index int) {
	/*
		This function takes as input a fasta file and it counts the number of nucleotide occurrences in this file
		and it returns a map for each nucleotide count...

		input_file: takes the fasta file path...
		flow: The connection channel through which the maps are flowing...flow
		index_flow: This channel is guidline for the master to know the data is from which routine...
		index: the id of the routine...
	*/

	in_file, err := os.Open(input_file)

	if err != nil {
		fmt.Println("Error in opening the file")
		return
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

	flow <- word_count
	index_flow <- index

}

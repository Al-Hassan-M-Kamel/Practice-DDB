package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){

	arg := flag.String("in_file", "", "The file path to process")

	flag.Parse()

	fmt.Println(Process_File(*arg))
}

func Process_File(file_path string) map[string]int{

	file, err := os.Open(file_path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	nt_counter := make(map[string]int)

	var line string

	reader.ReadString('\n') //skip the first line..

	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		for _, nt := range strings.TrimSpace(line) {
			nt_counter[string(nt)]++
		}
	}

	return  nt_counter

}

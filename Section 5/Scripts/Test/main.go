package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func CMD_Routine(file_path string, output_flow chan string) {
	output, _ := exec.Command("./Test/cmd/proc.exe", "-in_file", file_path).CombinedOutput()
	output_flow <- string(output)
}

func main() {

	files, _ := os.ReadDir("./Data")

	output_flow := make(chan string, len(files))

	for _, file := range files {
		file_path := fmt.Sprintf("./Data/%s", file.Name())

		go CMD_Routine(file_path, output_flow)

	}

	for i := 0; i < len(files); i++ {
		fmt.Println(<-output_flow)
	}

	// data_flow := make(chan map[string]int, len(files))

	// name_flow := make(chan string, len(files))

	// for _, file := range files {
	// 	file_path := fmt.Sprintf("./Data/%s", file.Name())
	// 	// go Process_File(file_path, data_flow, file.Name(), name_flow)
	// }

	// out_file , _ := os.Create("./out.txt")
	// defer out_file.Close()

	// writer := bufio.NewWriter(out_file)

	// for i:=0; i <len(files); i++{
	// 	// data := <- data_flow
	// 	// name := <- name_flow

	// 	fmt.Fprintf(writer, "The results of the file %s is :\n", name)

	// 	for k, v := range data{
	// 		fmt.Fprintf(writer, "%s: %d\n", k, v)
	// 	}

	// 	fmt.Fprintln(writer)
	// }

	// writer.Flush()

}

func Process_File(file_path string, data_flow chan map[string]int, file_name string, name_flow chan string) {

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

	data_flow <- nt_counter

	name_flow <- file_name

}

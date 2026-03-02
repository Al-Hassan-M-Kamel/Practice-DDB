package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

}

func Read_File_2() {
	file_path := "dataset_2_6.txt"

	file_handle, _ := os.Open(file_path)

	defer file_handle.Close()

	// reader := bufio.NewReader(file_handle)
	// buf, _ := io.ReadAll(reader)

	// Read the whole file one time...
	buf, _ := io.ReadAll(file_handle) // returns all the bytes in the file...

	// fmt.Printf("%s", buf)
	fmt.Println(string(buf))
}

func Read_File_1() {
	file_path := "dataset_2_6.txt"

	// The file_handle here is the pointer to our stream...
	file_handle, err := os.Open(file_path)

	if err != nil {
		fmt.Printf("An error occurred on opening the file \n")
		return
	}

	// Close the file...
	defer file_handle.Close() // This function will be executed in the end of the program...

	// now let's read the content in the file...
	reader := bufio.NewReader(file_handle)

	// We will read the file line by line...
	for {

		line, read_err := reader.ReadString('\n') // read till the new line...

		if read_err == io.EOF {
			// means we reach the end of the file...
			return
		}

		fmt.Println(line)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	in_file_path := "dataset_2_6.txt"
	out_file_path := "output.txt"

	// open file for reading...
	file1, _ := os.OpenFile(in_file_path, os.O_RDONLY, 0)

	file2, _ := os.OpenFile(out_file_path, os.O_WRONLY|os.O_CREATE, 0666)

	defer file1.Close()
	defer file2.Close()

	// read the content of the first file...

	bytes, _ := io.ReadAll(file1)
	data := string(bytes)

	// Write the content in the second file...

	writer := bufio.NewWriter(file2)
	writer.WriteString(data)

	writer.Flush()
}

func Read_File_3() {
	file_path := "dataset_2_6.txt"

	file_handle, _ := os.Open(file_path)

	defer file_handle.Close()

	scanner := bufio.NewScanner(file_handle)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
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

func Write_File_1() {
	/*
		os.OpenFile take three arguments:
		1- file path
		2- file mode one of (os.O_RDONLY, os.WRONLY, os.O_CREATE, os.O_TRUNC)
		3- the file permission
	*/
	file_handle, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}

	defer file_handle.Close()

	// create a writer to our stream...
	writer := bufio.NewWriter(file_handle)

	output_string := "Hello World...\n"

	// Here we write ten lines...
	for i := 0; i < 10; i++ {
		writer.WriteString(output_string)
	}

	writer.Flush() // The actual call that saves and writes all our contents and updates...
}

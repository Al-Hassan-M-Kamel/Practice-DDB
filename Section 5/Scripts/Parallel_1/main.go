package main

import (
	"fmt"
	"os"
	"os/exec"
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
}

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	

	// First of all set your command...

	output, err := exec.Command("../cmd/main.exe", "-in", "../Data/gene.fna").CombinedOutput()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))

}

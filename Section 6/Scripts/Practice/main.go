package main

import (
	"Scripts/Services/SServices"
	"fmt"
	"net/http"
)

func Save_Handler() {

	/*

		Implement the save function to save files coming from Master in a dir called Master_files
		located in the same dir with the script...
	*/

	// 1- Check Method Only POST allowed...

	// 2- Parse Multipart files with maximum 10 MB in RAM...

	// 3- Read the file content part part...

	// 4- Create folder to store files in located in the same dir with this script...

	// 5- Save the files in the new folder...

}

func main() {

	server := http.Server{
		Addr: fmt.Sprintf("%s:9080", SServices.Get_IP()),
	}

	http.HandleFunc("/save", Save_Handler)

	fmt.Println("Listen on: 9080")

	server.ListenAndServe()

}

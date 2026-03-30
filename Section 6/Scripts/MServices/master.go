package MServices

import (
	"bytes"
	"fmt"
	"io"
	"math/rand/v2"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

func Organize_Data(folder_path string, number_of_slaves int) map[string][]string {
	/*
		This function takes a folder of files only, better to be the same type and then returns
		a map contains the salve number as a key and its value are a slice of file paths.

		folder_path: the path of the folder in the source host to send its files...
		number_of_slaves: the number of nodes in the cluster for which files are sent...
	*/

	// 1- Read the Dir content...
	files, err := os.ReadDir(folder_path)

	if err != nil {
		fmt.Println("Error in reading folder Content...")
	}

	// 2- Get the file paths...
	file_paths := make([]string, 0)
	for _, f := range files {
		file_paths = append(file_paths, folder_path+"//"+f.Name())
	}

	// 3- Organize the files and distribute them across the slaves in the cluster...
	number_of_chunks := len(files) / number_of_slaves
	index := 0
	organizer := make(map[string][]string)

	for i := 0; i < number_of_slaves; i++ {
		index = i * number_of_chunks
		organizer[fmt.Sprintf("Slave: %d", i)] = file_paths[index : index+number_of_chunks]
	}

	// 4- Handle the remaining files...
	rand_slave := rand.IntN(number_of_slaves) // choose a random slave to sent the files to...
	organizer[fmt.Sprintf("Slave: %d", rand_slave)] = append(organizer[fmt.Sprintf("Slave: %d", rand_slave)], file_paths[number_of_chunks*number_of_slaves:]...)

	return organizer

}

func Map_Files(list_of_files []string, slave_url string, slave_id string, master_flow chan string) {

	fmt.Printf("Start Sending %d Files To: %s...\n", len(list_of_files), slave_id)

	flow := make(chan string, len(list_of_files))

	for _, file_path := range list_of_files {
		go Send_File(file_path, path.Base(file_path), fmt.Sprintf("%s/save", slave_url), flow)
	}

	for i := 0; i < len(list_of_files); i++ {
		<-flow
	}

	fmt.Printf("All File Are Sent TO %s Successfully With Total Number OF Files: %d\n", slave_id, len(list_of_files))

	master_flow <- "Done..."

}

func Send_File(file_path string, file_name string, url string, flow chan string) {
	/*
		This function responsible for sending one file throuth the network.
		file_path: the path of the file in the source host to send...
		file_name: the file name which will be used to create the file in the destination...
	*/

	// 1- Open the file for reading...
	file, err := os.Open(file_path)

	if err != nil {
		fmt.Println("Error in Opening the file...")
		return
	}

	defer file.Close()

	// 2- Declare the bytes buffer in which the file content will be read...

	var body bytes.Buffer

	// 3- Define the writer to write the message body...

	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", file_name)

	if err != nil {
		fmt.Println("Error in creating the form content...")
		return
	}

	// 4- Write the file content into the message body...
	io.Copy(part, file)
	writer.Close()

	// 5- Create the http request and send the data...
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		fmt.Println("Error in creating a new http request...")
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error in sending the data...")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		flow <- "The file Sent Sucessfully..."
	}

}

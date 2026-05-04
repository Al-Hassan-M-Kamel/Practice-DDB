package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"
)

type Slave struct {
	IP       string `json:"ip"`
	Slave_id int    `json:"slave"`
}

func Config(config_file_path string) map[int]string {

	json_file, err := os.Open(config_file_path)
	if err != nil {
		fmt.Println("Error in opening file")
		return nil
	}

	json_data, err := io.ReadAll(json_file)
	if err != nil {
		fmt.Println("Error in reading json content")
		return nil
	}

	var slaves []Slave
	json.Unmarshal(json_data, &slaves)
	slaves_mp := make(map[int]string, len(slaves))
	for i := 0; i < len(slaves); i++ {
		slaves_mp[slaves[i].Slave_id] = slaves[i].IP
	}

	return slaves_mp

}

func Organize_Dir_Data(dir_path string, slaves_count int) map[int][]string {

	/*
		This function takes a folder of files only, better to be the same type and then returns
		a map contains the salve number as a key and its value are a slice of file paths.

		folder_path: the path of the folder in the source host to send its files...
		number_of_slaves: the number of nodes in the cluster for which files are sent...
	*/

	// 1- Read the Dir content...
	files, err := os.ReadDir(dir_path)
	if err != nil {
		fmt.Println("Error in reading folder content...")
		return nil
	}

	// 2- Get the file paths...
	file_paths := make([]string, 0)
	for _, file := range files {
		file_paths = append(file_paths, dir_path+"//"+file.Name())
	}

	// 3- Organize the files and distribute them across the slaves in the cluster...
	number_of_chunks := len(files) / slaves_count
	index := 0
	organizer := make(map[int][]string)
	for i := 0; i < slaves_count; i++ {
		index = i * number_of_chunks
		organizer[i+1] = file_paths[index : index+number_of_chunks]
	}

	// 4- Handle the remaining files...
	rand_slave := rand.IntN(slaves_count) + 1 // choose a random slave to sent the files to...
	organizer[rand_slave] = append(organizer[rand_slave], file_paths[number_of_chunks*slaves_count:]...)

	return organizer

}

func Send_File(file_path string, file_name string, url string, flow chan<- int) {
	/*
		This function responsible for sending one file throuth the network.
		file_path: the path of the file in the source host to send...
		file_name: the file name which will be used to create the file in the destination...
		url: the destination url like: http://127.0.0.1:9000/endPointName ...
		flow: a channel used for synchronization...
	*/

	// 1- Open the file for reading...
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error in Opening the file...")
		flow <- 0
		return
	}

	defer file.Close()

	// 2- Declare the bytes buffer in which the file content will be read...
	var file_body bytes.Buffer

	writer := multipart.NewWriter(&file_body)
	part, err := writer.CreateFormFile("file", file_name)
	if err != nil {
		fmt.Println("Error in creating the form content...")
		return
	}

	// 4- Write the file content into the message body...
	io.Copy(part, file)
	writer.Close()

	// 5- Create the http request and send the data...
	req, err := http.NewRequest("Post", url, &file_body)
	if err != nil {
		fmt.Println("Error in creating a new http request...")
		flow <- 0
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		// fmt.Println("Error in sending the data...")
		flow <- 0
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		flow <- 1
	}

}

func Map_Files(list_of_files []string, slave_url string, slave_id int, master_flow chan<- string) {
	fmt.Printf("Start Sending %d Files To Slave # %d ... \n", len(list_of_files), slave_id)

	flow := make(chan int, len(list_of_files))

	for _, file_path := range list_of_files {
		go Send_File(file_path, path.Base(file_path), slave_url, flow)
	}

	sent_files := 0
	for i := 0; i < len(list_of_files); i++ {
		sent_files += <-flow
	}

	fmt.Printf("%d Files Are Completely Sent to Slave #%d\n", sent_files, slave_id)
	master_flow <- fmt.Sprintf("Done Sending for Slave # %d...\n", slave_id)
}

func Parse_Slave_URL(slave_ip string, port_number string, end_point string) string {
	return fmt.Sprintf("http://%s:%s/%s", slave_ip, port_number, end_point)
}

func main() {

	nodes_config := Config("./MServices/config.json")

	dir_organizer := Organize_Dir_Data("Data", len(nodes_config))
	master_flow := make(chan string, len(nodes_config))

	for node_id, files := range dir_organizer {
		node_url := Parse_Slave_URL(nodes_config[node_id], "5000", "save")
		go Map_Files(files, node_url, node_id, master_flow)
	}

	for i := 0; i < len(nodes_config); i++ {
		fmt.Println(<-master_flow)
	}

}

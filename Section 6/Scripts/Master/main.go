package main

import (
	"Scripts/Services/MServices"
	"fmt"
)

func main() {

	folder_path := "./Data"
	number_of_slaves := 3

	// 1- Get IPs for slaves...
	slave_ips := make(map[string]string)
	for i := 0; i < 3; i++ {
		slave_ips[fmt.Sprintf("Slave: %d", i)] = fmt.Sprintf("http://127.0.0.1:908%d", i)
	}

	data := MServices.Organize_Data(folder_path, number_of_slaves)

	master_flow := make(chan string, number_of_slaves)

	for k, v := range data {
		go MServices.Map_Files(v, slave_ips[k], k, master_flow)
	}

	for i := 0; i < number_of_slaves; i++ {
		<-master_flow
	}

	fmt.Println("Done Sending...")
}

package main

import (
	"Cluster/Services/MServices"
	"fmt"
	"path"
)

func main() {

	// slaves_config := MServices.Config()
	url := MServices.Parse_Slave_URL("127.0.0.1", "9080", "save")
	fmt.Println(url)

	MServices.Send_File("F:\\Work\\Courses\\DDB\\Material\\Section 6\\Scripts\\Data\\gene - Copy - Copy - Copy.fna", path.Base("./Data/gene - Copy - Copy - Copy.fna"), url)

}

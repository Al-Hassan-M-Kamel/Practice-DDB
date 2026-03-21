package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main3() {

	file_path := "F:\\Work\\Courses\\DDB\\Material\\Section 5\\Scripts\\Services\\ServiceA\\gene.fna"

	file, err := os.Open(file_path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", file_path)
	if err != nil {
		panic(err)
	}

	io.Copy(part, file)

	writer.Close()

	req, err := http.NewRequest("POST", "http://127.0.0.1:9080/file", &body)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("File Uploaded Successfully...")

}

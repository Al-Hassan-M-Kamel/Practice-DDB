package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetFile(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) //Up to 10 MB is stored in RAM other goes to temp files...
	if err != nil {
		http.Error(w, "File is too large", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error in Retreiving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println("File Received: ", handler.Filename)

	dst, err := os.Create("./uploaded" + handler.Filename)

	if err != nil {
		http.Error(w, "Error in Saving file", http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	io.Copy(dst, file)

	w.Write([]byte("File uploaded Successfully..."))
}

func main3() {

	server := http.Server{
		Addr: "127.0.0.1:9080",
	}

	fmt.Println("Listen on :9080")

	http.HandleFunc("/file", GetFile)

	server.ListenAndServe()
}

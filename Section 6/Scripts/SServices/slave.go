package SServices

import (
	"io"
	"net"
	"net/http"
	"os"
)

func Get_IP() string {

	/*
		This function returns the IP address of the current local host...
	*/

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Skip loopback and IPv6 addresses
			if ip.IsLoopback() || ip.To4() == nil {
				continue
			}

			return ip.String()
		}
	}
	return ""
}

func Save_Handler(w http.ResponseWriter, r *http.Request) {

	// 1- Check Method Only POST allowed...
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2- Parse Multipart files with maximum 10 MB in RAM...
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File is too Large", http.StatusBadRequest)
		return
	}

	// 3- Read the file content part part...
	file, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Error in Retreive file content...", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// 4- Create folder to store files in located in the same dir with this script...
	err = os.MkdirAll("./Slave/Master_Files", os.ModePerm)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 5- Save the files in the new folder...

	dst, err := os.Create("./Slave/Master_Files/" + handler.Filename)
	if err != nil {
		http.Error(w, "Error in Saving file", http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	io.Copy(dst, file)

}

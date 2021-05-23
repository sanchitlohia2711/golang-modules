package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	createImageHandler := http.HandlerFunc(createImage)
	http.Handle("/image", createImageHandler)
	http.ListenAndServe(":8080", nil)
}

func createImage(w http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	file, h, err := request.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tmpfile, err := os.Create("./" + h.Filename)
	defer tmpfile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	return
}

// curl --location --request PUT 'http://localhost:8080/employee' \
// --header 'Content-Type: multipart/form-data' \
// --form 'name=John' \
// --form 'age=23' \
// --form 'photo=@test.png'

// curl --location --request PUT 'http://localhost:8080/employee' \
// --header 'Content-Type: multipart/form-data' \
// --form 'name=John' \
// --form 'age=23' \
// --form 'photo=@test1.png' \
// --form 'photo=@test2.png'

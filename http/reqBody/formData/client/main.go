package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	call("http://localhost:8080/employee", "POST")
}

func call(urlPath, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormField("name")
	if err != nil {

	}
	_, err = io.Copy(fw, strings.NewReader("John"))
	if err != nil {
		return err
	}

	fw, err = writer.CreateFormField("age")
	if err != nil {

	}
	_, err = io.Copy(fw, strings.NewReader("23"))
	if err != nil {
		return err
	}

	fw, err = writer.CreateFormFile("photo", "test.png")
	if err != nil {

	}
	file, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}

	// Close multipart writer.
	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:8080/employee", bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	return nil
}

//curl -v -X POST http://localhost:8085/example -H "content-type: application/json" -H "foo: bar1" -H "foo: bar2"

//curl -v -X GET http://localhost:8085/employee -d '{"Name":"s"}'

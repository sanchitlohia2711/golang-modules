package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type employee struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Photo string `json:"photo"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/employee?name=John")
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}
	if resp.StatusCode == http.StatusNotFound {
		log.Fatalf("Employee with name John not found. Please create employee first")
	}
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", string(body))
}

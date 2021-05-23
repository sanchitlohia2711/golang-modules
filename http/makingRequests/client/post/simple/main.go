package main

import (
	"bytes"
	"encoding/json"
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
	e := employee{
		Name: "John",
		Age:  21,
	}

	b, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}

	resp, err := http.Post("http://localhost:8080/employee", "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}
	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", string(body))
}

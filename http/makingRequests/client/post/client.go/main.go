package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client http.Client

func init() {
	client = http.Client{
		Timeout: time.Second * 5,
	}
}

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

	req, err := http.NewRequest("POST", "http://localhost:8080/employee", bytes.NewBuffer(b))
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
	}
	req.Header.Add("foo", "bar")
	resp, err := client.Do(req)
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

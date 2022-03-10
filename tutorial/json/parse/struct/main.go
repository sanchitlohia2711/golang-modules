package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type employee struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {

	var emp employee
	file, err := ioutil.ReadFile("employee.json")
	if err != nil {
		log.Fatalf("Some error occured while reading file. Error: %s", err)
	}
	err = json.Unmarshal([]byte(file), &emp)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Printf("emp Struct: %#v\n", emp)
}

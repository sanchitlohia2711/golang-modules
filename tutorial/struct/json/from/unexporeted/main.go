package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee1 struct {
	Name   string `json:"n"`
	Age    int    `json:"a"`
	salary int    `json:"s"`
}

func main() {
	e1Json := `{"n":"John","a":21,"salary":1000}`

	var e1Converted employee1
	err := json.Unmarshal([]byte(e1Json), &e1Converted)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Printf("employee1 Struct: %#v\n", e1Converted)

}

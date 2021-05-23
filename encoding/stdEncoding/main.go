package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	sample := "a@"
	encodedString := base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Println(encodedString)

	originalStringBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))

	sample = "ab@c"
	encodedString = base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Println(encodedString)

	originalStringBytes, err = base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))

	sample = "ab@cd"
	encodedString = base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Println(encodedString)

	originalStringBytes, err = base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))
}

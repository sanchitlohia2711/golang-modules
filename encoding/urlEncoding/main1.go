package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func maina() {
	sampl := []byte{252, 123, 189}

	fmt.Println(string(sampl))
	encodedString := base64.StdEncoding.EncodeToString([]byte(sampl))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(sampl)))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte(sampl)))

	fmt.Println(base64.StdEncoding.EncodeToString([]byte("�")))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("�")))

	originalStringBytes, err := base64.URLEncoding.DecodeString("a/b=")
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))

	sample := "ab@c"
	encodedString = base64.URLEncoding.EncodeToString([]byte(sample))
	fmt.Println(encodedString)

	originalStringBytes, err = base64.URLEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))

	sample = "ab@cd"
	encodedString = base64.URLEncoding.EncodeToString([]byte(sample))
	fmt.Println(encodedString)

	originalStringBytes, err = base64.URLEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))
}

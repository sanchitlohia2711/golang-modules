package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	input_url := "https://test:abcd123@golangbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"
	u, err := url.Parse(input_url)
	if err != nil {
		log.Fatal(err)
	}

	fullHostname := u.Scheme + "://" + u.Host

	fmt.Println(fullHostname)
}

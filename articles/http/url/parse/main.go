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

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Hostname())
	fmt.Println(u.Port())
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)
	fmt.Println(u.String())
}

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileBytes, err := ioutil.ReadFile("test.png")
	if err != nil {
		panic(err)
	}

	fileString := string(fileBytes)

	fmt.Println(fileString)
}

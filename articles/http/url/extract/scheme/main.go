package main

import (
	"fmt"
	"log"

	"mvdan.cc/xurls/v2"
)

func main() {
	xurlsStrict, err := xurls.StrictMatchingScheme("https")
	if err != nil {
		log.Fatalf("Some error occured. Error: %s", err)
	}
	input := "The webiste is https://golangbyexample.com:8000/tutorials/intro amd mail to mailto:contactus@golangbyexample.com"
	output := xurlsStrict.FindAllString(input, -1)
	fmt.Println(output)
}

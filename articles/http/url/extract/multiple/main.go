package main

import (
	"fmt"

	"mvdan.cc/xurls/v2"
)

func main() {
	xurlsStrict := xurls.Strict()
	input := "The webiste is https://golangbyexample.com:8000/tutorials/intro amd mail to mailto:contactus@golangbyexample.com"
	output := xurlsStrict.FindAllString(input, -1)
	fmt.Println(output)
}

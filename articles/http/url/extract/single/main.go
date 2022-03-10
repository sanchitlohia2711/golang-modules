package main

import (
	"fmt"

	"mvdan.cc/xurls/v2"
)

func main() {
	xurlsStrict := xurls.Strict()
	output := xurlsStrict.FindAllString("golangbyexample.com is https://golangbyexample.com", -1)
	fmt.Println(output)

	xurlsRelaxed := xurls.Relaxed()
	output = xurlsRelaxed.FindAllString("The website is golangbyexample.com", -1)
	fmt.Println(output)

	output = xurlsRelaxed.FindAllString("golangbyexample.com is https://golangbyexample.com", -1)
	fmt.Println(output)
}

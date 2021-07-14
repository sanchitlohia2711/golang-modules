package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`(\w+):\1`)

	input := "The names are John:John"

	result := sampleRegexp.ReplaceAllString(input, "$1")
	fmt.Println(string(result))
}

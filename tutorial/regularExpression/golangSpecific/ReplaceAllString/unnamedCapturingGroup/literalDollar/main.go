package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`(\w+):([0-9]\d{1,3})`)

	input := "The names are John:21, Simon:23, and Mike:19"

	result := sampleRegexp.ReplaceAllLiteralString(input, "$1")
	fmt.Println(string(result))
}

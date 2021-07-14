package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`\w+:[0-9]\d{1,3}`)

	input := "The persons are John:21, Simon:23, Mike:19"

	result := sampleRegexp.ReplaceAllString(input, "{redacted}")
	fmt.Println(string(result))
}

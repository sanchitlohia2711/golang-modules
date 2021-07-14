package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("^([A-Z][a-z]+):([0-9]\\d{1,3})$")

	matches := sampleRegexp.FindStringSubmatch("John:21")
	fmt.Println(matches)
}

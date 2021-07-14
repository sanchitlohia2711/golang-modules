package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("^[A-Z][a-z]+:[0-9]\\d{0,3}$")

	match := sampleRegexp.Match([]byte("John:21"))
	fmt.Printf("For John:21: %t\n", match)
}

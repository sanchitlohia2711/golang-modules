package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("a\\.b")

	match := sampleRegexp.Match([]byte("a.b"))

	fmt.Printf("For a.b string: %t\n", match)
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile("\\^abc")

	match := sampleRegex.Match([]byte("^abc"))
	fmt.Printf("For ^abc: %t\n", match)

	match = sampleRegex.Match([]byte("abc"))
	fmt.Printf("For abc: %t\n", match)
}

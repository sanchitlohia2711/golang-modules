package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile("(?m)abc$")

	match := sampleRegex.Match([]byte("1abc\n1abd\n1abd"))
	fmt.Printf("For 1abc\\n1abd\\n1abd: %t\n", match)

	match = sampleRegex.Match([]byte("1abd\n1abc\n1abd"))
	fmt.Printf("For 1abd\\n1abc\\n1abd: %t\n", match)

	match = sampleRegex.Match([]byte("1abd\n1abd\n1abc"))
	fmt.Printf("For 1abd\\n1abd\\n1abc: %t\n", match)

	match = sampleRegex.Match([]byte("abd\nabd\nabd"))
	fmt.Printf("For 1abd\\n1abd\\n1abd: %t\n", match)
}

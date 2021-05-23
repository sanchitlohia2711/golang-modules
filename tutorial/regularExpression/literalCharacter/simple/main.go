package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile("abc")

	match := sampleRegex.Match([]byte("abc"))
	fmt.Printf("For abc: %t\n", match)

	match = sampleRegex.Match([]byte("1abc2"))
	fmt.Printf("For 1abc2: %t\n", match)

	match = sampleRegex.Match([]byte("xyz"))
	fmt.Printf("For xyz: %t\n", match)
}

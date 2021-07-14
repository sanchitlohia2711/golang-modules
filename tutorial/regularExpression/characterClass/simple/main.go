package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile(`a[bc]d`)

	match := sampleRegex.MatchString("abd")
	fmt.Printf("For abd: %t\n", match)

	match = sampleRegex.MatchString("acd")
	fmt.Printf("For acd: %t\n", match)

	match = sampleRegex.MatchString("abc")
	fmt.Printf("For abc: %t\n", match)

	match = sampleRegex.MatchString("abcd")
	fmt.Printf("For abcd: %t\n", match)
}

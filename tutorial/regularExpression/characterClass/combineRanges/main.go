package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile(`^[0-9a-zA-Z]$`)

	match := sampleRegex.MatchString("9")
	fmt.Printf("For 9: %t\n", match)

	match = sampleRegex.MatchString("A")
	fmt.Printf("For A: %t\n", match)

	match = sampleRegex.MatchString("y")
	fmt.Printf("For y: %t\n", match)

	match = sampleRegex.MatchString("a9")
	fmt.Printf("\nFor y: %t\n", match)
}

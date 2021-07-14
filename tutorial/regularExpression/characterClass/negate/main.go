package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile(`^[^ab]$`)

	match := sampleRegex.MatchString("a")
	fmt.Printf("For a: %t\n", match)

	match = sampleRegex.MatchString("b")
	fmt.Printf("For b: %t\n", match)

	match = sampleRegex.MatchString("d")
	fmt.Printf("\nFor d: %t\n", match)

	match = sampleRegex.MatchString("1")
	fmt.Printf("For 1: %t\n", match)
}

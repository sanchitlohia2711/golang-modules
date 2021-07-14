package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile(`a[123bc]d`)

	match := sampleRegex.MatchString("a1d")
	fmt.Printf("For a1d: %t\n", match)

	match = sampleRegex.MatchString("a2d")
	fmt.Printf("For a2d: %t\n", match)

	match = sampleRegex.MatchString("a3d")
	fmt.Printf("For a3d: %t\n", match)

	match = sampleRegex.MatchString("abd")
	fmt.Printf("For abd: %t\n", match)

	match = sampleRegex.MatchString("acd")
	fmt.Printf("For acd: %t\n", match)

	match = sampleRegex.MatchString("a12d")
	fmt.Printf("\nFor a12d: %t\n", match)

	match = sampleRegex.MatchString("a23d")
	fmt.Printf("For a23d: %t\n", match)
}

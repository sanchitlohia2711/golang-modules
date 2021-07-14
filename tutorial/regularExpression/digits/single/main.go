package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`\d`)

	fmt.Println("For regex \\d")
	match := sampleRegexp.MatchString("1")
	fmt.Printf("For 1: %t\n", match)

	match = sampleRegexp.MatchString("4")
	fmt.Printf("For 4: %t\n", match)

	match = sampleRegexp.MatchString("9")
	fmt.Printf("For 9: %t\n", match)

	match = sampleRegexp.MatchString("a")
	fmt.Printf("For a: %t\n", match)

	sampleRegexp = regexp.MustCompile(`5`)
	fmt.Println("\nFor regex 5")
	match = sampleRegexp.MatchString("5")
	fmt.Printf("For 5: %t\n", match)

	match = sampleRegexp.MatchString("6")
	fmt.Printf("For 6: %t\n", match)
}

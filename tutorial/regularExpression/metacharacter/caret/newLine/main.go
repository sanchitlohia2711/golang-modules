package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile("(?m)^abc")

	match := sampleRegex.Match([]byte("abc\nabd\nabd"))
	fmt.Printf("For abc\\nabd\\nabd: %t\n", match)

	match = sampleRegex.Match([]byte("abd\nabc\nabd"))
	fmt.Printf("For abd\\nabc\\nabd: %t\n", match)

	match = sampleRegex.Match([]byte("abd\nabd\nabc"))
	fmt.Printf("For abd\\nabd\\nabc: %t\n", match)

	match = sampleRegex.Match([]byte("abd\nabd\nabd"))
	fmt.Printf("For abd\\nabd\\nabd: %t\n", match)

	match = sampleRegex.Match([]byte("abc"))
	fmt.Printf("For abc: %t\n", match)

	match = sampleRegex.Match([]byte("1abc23"))
	fmt.Printf("For 1abc23: %t\n", match)
}

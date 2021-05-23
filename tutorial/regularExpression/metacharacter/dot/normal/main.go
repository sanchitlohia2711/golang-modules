package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(".")

	match := sampleRegexp.Match([]byte("a"))
	fmt.Printf("For a: %t\n", match)

	match = sampleRegexp.Match([]byte("b"))
	fmt.Printf("For b: %t\n", match)

	match = sampleRegexp.Match([]byte("ab"))
	fmt.Printf("For ab: %t\n", match)

	match = sampleRegexp.Match([]byte(""))
	fmt.Printf("For empty string: %t\n", match)
}

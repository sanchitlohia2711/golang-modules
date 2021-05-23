package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("(?s).")

	match := sampleRegexp.Match([]byte("\n"))
	fmt.Printf("For \\n: %t\n", match)
}

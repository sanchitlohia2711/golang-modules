package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("\\(xyz\\)")

	match := sampleRegexp.Match([]byte("(xyz)"))
	fmt.Printf("For (xyz): %t\n", match)

	sampleRegexp = regexp.MustCompile(`\(xyz\)`)

	match = sampleRegexp.Match([]byte("(xyz)"))
	fmt.Printf("For (xyz): %t\n", match)
}

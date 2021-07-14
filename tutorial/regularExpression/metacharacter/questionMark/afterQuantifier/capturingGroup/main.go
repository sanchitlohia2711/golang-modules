package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("(a+?)(a*)")

	match := sampleRegexp.FindStringSubmatch("aaaaaaa")
	fmt.Printf("Match: %s Length: %d\n", match, len(match))

	sampleRegexp = regexp.MustCompile("(a*?)(a*)")

	match = sampleRegexp.FindStringSubmatch("aaaaaaa")
	fmt.Printf("Match: %s Length: %d\n", match, len(match))
}

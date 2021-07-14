package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("http(s+?)")

	match := sampleRegexp.Find([]byte("Better is httpsss"))
	fmt.Printf("Match: %s\n", match)

	sampleRegexp = regexp.MustCompile("http(s*?)")

	match = sampleRegexp.Find([]byte("Better is httpsss"))
	fmt.Printf("Match: %s\n", match)
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("https?")

	match := sampleRegexp.Find([]byte("Better is https"))
	fmt.Printf("Match: %s\n", match)
}

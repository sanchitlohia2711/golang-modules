package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`ab{2,4}?`)

	matches := sampleRegexp.FindString("abb")
	fmt.Println(matches)

	matches = sampleRegexp.FindString("abbb")
	fmt.Println(matches)

	matches = sampleRegexp.FindString("abbbb")
	fmt.Println(matches)
}

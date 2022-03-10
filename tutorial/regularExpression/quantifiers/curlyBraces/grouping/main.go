package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`(ab){2}`)

	matches := sampleRegexp.FindString("abab")
	fmt.Println(matches)

	matches = sampleRegexp.FindString("ababbc")
	fmt.Println(matches)
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`b{2}?`)

	matches := sampleRegexp.FindStringSubmatch("bb")
	fmt.Println(matches)

	matches = sampleRegexp.FindStringSubmatch("bbb")
	fmt.Println(matches)

	matches = sampleRegexp.FindStringSubmatch("bbbb")
	fmt.Println(matches)
}

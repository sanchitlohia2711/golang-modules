package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("^abc(de)?$")

	match := sampleRegexp.Match([]byte("abc"))
	fmt.Printf("For abc: %t\n", match)

	match = sampleRegexp.Match([]byte("abcde"))
	fmt.Printf("For abcde: %t\n", match)

	match = sampleRegexp.Match([]byte("abcd"))
	fmt.Printf("For abcd: %t\n", match)
}

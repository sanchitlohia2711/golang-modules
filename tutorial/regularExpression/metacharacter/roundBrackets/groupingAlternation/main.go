package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Regex s(am)+")
	sampleRegexp := regexp.MustCompile("s(am)?")

	match := sampleRegexp.Match([]byte("sam"))
	fmt.Printf("For sam: %t\n", match)

	match = sampleRegexp.Match([]byte("s"))
	fmt.Printf("For s: %t\n", match)
}

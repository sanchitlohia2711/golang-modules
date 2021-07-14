package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile(`\w*(?=is the name)`)

	match := sampleRegex.Match([]byte("John is the name"))
	fmt.Printf("Match: %t\n", match)

}

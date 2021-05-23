package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("[.]")
	match := sampleRegexp.Match([]byte("."))

	fmt.Println(match)

}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := `b+`
	sampleRegexp := regexp.MustCompile("a" + regex)

	match := sampleRegexp.FindString("abb")
	fmt.Println(match)

}

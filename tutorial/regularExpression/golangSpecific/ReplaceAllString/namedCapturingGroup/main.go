package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`(?P<Name>\w+):(?P<Age>[0-9]\d{1,3})`)

	input := "The names are John:21, Simon:23, Mike:19"

	result := sampleRegexp.ReplaceAllString(input, "$Age:$Name")
	fmt.Println(string(result))
}

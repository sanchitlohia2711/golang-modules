package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`(\w+):([0-9]\d{1,3})`)

	//sampleRegexp := regexp.MustCompile("^([A-Z][a-z]+):([0-9]\\d{1,3})$")
	//matches := sampleRegexp.FindStringSubmatch("John:21")
	matches := sampleRegexp.FindAllStringSubmatch("The names are John:21, Simon:23, Mike:19", -1)
	fmt.Println(matches)
}

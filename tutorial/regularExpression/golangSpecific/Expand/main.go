package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`(?P<name>\w+):(?P<age>[0-9]\d{1,3})`)

	//sampleRegexp := regexp.MustCompile("^([A-Z][a-z]+):([0-9]\\d{1,3})$")
	//matches := sampleRegexp.FindStringSubmatch("John:21")

	result := []byte{}
	sample := "The names are John:21, Simon:23, Mike:19"

	template := "$name=$age\n"

	for _, submatches := range sampleRegexp.FindAllStringSubmatchIndex(sample, -1) {
		result = sampleRegexp.ExpandString(result, template, sample, submatches)
	}
	fmt.Println(string(result))
}

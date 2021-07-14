package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile("^[A-Z][a-z]+-[1-9]\\d{1,3}$")

	match := sampleRegexp.Match([]byte("John-21"))
	fmt.Printf("For John-21: %t\n", match)

	group := sampleRegexp.FindString("John-21")
	fmt.Printf("For John-21: %s\n", group)

	sampleRegexp = regexp.MustCompile("^([A-Z][a-z]+)-[0-9]\\d{1,3}$")

	group = sampleRegexp.FindString("John-21")
	fmt.Printf("For John-21: %s\n", group)

	group2 := sampleRegexp.FindStringSubmatch("John-21")
	for _, part := range group2 {
		fmt.Println(string(part))
	}
}

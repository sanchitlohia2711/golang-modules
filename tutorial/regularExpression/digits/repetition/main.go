package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`\d+`)

	fmt.Println(`For regex \d+`)
	match := sampleRegexp.MatchString("12345")
	fmt.Printf("For 12345: %t\n", match)

	match = sampleRegexp.MatchString("")
	fmt.Printf("For empty string: %t\n", match)

	sampleRegexp = regexp.MustCompile(`\d*`)

	fmt.Println()
	fmt.Println(`For regex \d*`)
	match = sampleRegexp.MatchString("12345")
	fmt.Printf("For 12345: %t\n", match)

	match = sampleRegexp.MatchString("")
	fmt.Printf("For empty string: %t\n", match)

	sampleRegexp = regexp.MustCompile(`\d{2}`)

	fmt.Println()
	fmt.Println(`For regex \d{2}`)
	match = sampleRegexp.MatchString("12")
	fmt.Printf("For 12: %t\n", match)

	match = sampleRegexp.MatchString("1")
	fmt.Printf("For 1: %t\n", match)

}

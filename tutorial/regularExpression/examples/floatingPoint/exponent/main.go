package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`)

	fmt.Println("Valid Inputs")
	match := sampleRegex.MatchString("1.2")
	fmt.Printf("For 1.2: %t\n", match)

	match = sampleRegex.MatchString(".12")
	fmt.Printf("For .12: %t\n", match)

	match = sampleRegex.MatchString("12")
	fmt.Printf("For 12: %t\n", match)

	match = sampleRegex.MatchString("12.")
	fmt.Printf("For 12.: %t\n", match)

	match = sampleRegex.MatchString("+1.2")
	fmt.Printf("For +1.2.: %t\n", match)

	match = sampleRegex.MatchString("-1.2")
	fmt.Printf("For -1.2.: %t\n", match)

	match = sampleRegex.MatchString("1.2e3")
	fmt.Printf("For 1.2e3.: %t\n", match)

	fmt.Println("\nInValid Inputs")
	match = sampleRegex.MatchString(".")
	fmt.Printf("For .: %t\n", match)

	match = sampleRegex.MatchString("")
	fmt.Printf("For empty string: %t\n", match)

	match = sampleRegex.MatchString("00.1")
	fmt.Printf("For 00.1: %t\n", match)

	match = sampleRegex.MatchString("001")
	fmt.Printf("For 001 %t\n", match)

	match = sampleRegex.MatchString("+")
	fmt.Printf("For +: %t\n", match)

	match = sampleRegex.MatchString("-")
	fmt.Printf("For -: %t\n", match)

	match = sampleRegex.MatchString("+.")
	fmt.Printf("For +.: %t\n", match)

	match = sampleRegex.MatchString("-.")
	fmt.Printf("For -.: %t\n", match)

	match = sampleRegex.MatchString("1.e2")
	fmt.Printf("For 1.e2: %t\n", match)

	match = sampleRegex.MatchString(".e2")
	fmt.Printf("For .e2: %t\n", match)

	match = sampleRegex.MatchString("a1.2")
	fmt.Printf("For a1.2 %t\n", match)

	match = sampleRegex.MatchString("1.2b")
	fmt.Printf("For 1.2b %t\n", match)

	match = sampleRegex.MatchString("a1.2b")
	fmt.Printf("For a1.2b %t\n", match)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	output := isPalindrome(121)
	fmt.Println(output)

	output = isPalindrome(12)
	fmt.Println(output)

	output = isPalindrome(1234)
	fmt.Println(output)

	output = isPalindrome(12321)
	fmt.Println(output)

	output = isPalindrome(-101)
	fmt.Println(output)

}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	xReversed := reverse(x)

	return xReversed == x
}

func reverse(x int) int {
	sign := "positive"
	if x >= 0 {
		sign = "positive"
	} else {
		sign = "negative"
	}

	x = int(math.Abs(float64(x)))

	var reversedDigit int

	for x > 0 {
		lastDigit := x % 10
		reversedDigit = reversedDigit*10 + lastDigit

		x = x / 10
	}

	if sign == "negative" {
		reversedDigit = reversedDigit * -1
	}

	if reversedDigit < int(math.Pow(2, 31)*-1) || reversedDigit > int(math.Pow(2, 31)-1) {
		return 0
	}
	return reversedDigit
}

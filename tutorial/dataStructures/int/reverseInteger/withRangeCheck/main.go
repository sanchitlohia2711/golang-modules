package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	reversedInteger := reverse(123)

	fmt.Println(reversedInteger)
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

	var a int
	intSizeInBytes := unsafe.Sizeof(a)
	var maxPower float64
	if intSizeInBytes == 8 {
		maxPower = 63
	} else {
		maxPower = 31
	}
	if reversedDigit < int(math.Pow(2, maxPower)*-1) {
		fmt.Println("jhere")
	}

	p := int(math.Pow(2, maxPower) - 1)
	fmt.Println(p)
	if reversedDigit > int(math.Pow(2, maxPower)-1) {
		fmt.Println("jhere")
	}

	if reversedDigit < int(math.Pow(2, maxPower)*-1) || reversedDigit > int(math.Pow(2, maxPower)-1) {
		return 0
	}
	return reversedDigit
}

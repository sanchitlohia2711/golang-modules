package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {

	sign := 1
	if dividend < 0 || divisor < 0 {
		sign = -1
	}

	if dividend < 0 && divisor < 0 {
		sign = 1
	}

	if dividend < 0 {
		dividend = -1 * dividend
	}

	if divisor < 0 {
		divisor = -1 * divisor
	}

	start := divisor

	i := 0

	for start <= dividend {
		start = start + divisor
		i++
	}

	output := i * sign

	if output > int((math.Pow(2, 31) - 1)) {
		output = int(math.Pow(2, 31) - 1)
	}

	if output < int(math.Pow(-2, 31)) {
		output = int(math.Pow(-2, 31))
	}

	return output
}

func main() {
	output := divide(15, 2)
	fmt.Println(output)

	output = divide(-15, 2)
	fmt.Println(output)

	output = divide(15, -2)
	fmt.Println(output)

	output = divide(-15, -2)
	fmt.Println(output)
}

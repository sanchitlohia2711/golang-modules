package main

import "fmt"

func plusOne(digits []int) []int {

	lenDigits := len(digits)

	output := make([]int, 0)

	lastDigit := digits[lenDigits-1]

	add := lastDigit + 1
	carry := add / 10

	lastDigit = add % 10
	output = append(output, lastDigit)

	for i := lenDigits - 2; i >= 0; i-- {
		o := digits[i] + carry
		lastDigit = o % 10
		carry = o / 10
		output = append(output, lastDigit)
	}

	if carry == 1 {
		output = append(output, 1)
	}

	return reverse(output, len(output))
}

func reverse(input []int, length int) []int {
	start := 0
	end := length - 1

	for start < end {
		input[start], input[end] = input[end], input[start]
		start++
		end--
	}

	return input
}

func main() {
	output := plusOne([]int{1, 2})
	fmt.Println(output)

	output = plusOne([]int{9, 9})
	fmt.Println(output)
}

package main

import "fmt"

func addDigits(num int) int {

	if num < 10 {
		return num
	}

	for num > 9 {
		num = sum(num)
	}

	return num

}

func sum(num int) int {
	output := 0
	for num > 0 {
		output = output + num%10
		num = num / 10
	}
	return output
}

func main() {
	output := addDigits(453)
	fmt.Println(output)

	output = addDigits(45)
	fmt.Println(output)

}

package main

import "fmt"

func main() {
	// a := 121
	// output := isPalindrome(a, &a)
	// fmt.Println(output)

	// a = 12
	// output = isPalindrome(a, &a)
	// fmt.Println(output)

	// a = 1234
	// output = isPalindrome(a, &a)
	// fmt.Println(output)

	a := 12321
	output := isPalindrome(a, &a)
	fmt.Println(output)

	a = -121
	output = isPalindrome(-a, &a)
	fmt.Println(output)

}

func isPalindrome(x int, dup *int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	palin := isPalindrome(x/10, dup)

	*dup = *dup / 10
	lastDigit := x % 10

	if palin && *dup%10 == lastDigit {
		return true
	}

	return false

}

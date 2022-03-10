package main

import (
	"fmt"
	"strconv"
)

func main() {
	output := myAtoi("121")
	fmt.Println(output)

	output = myAtoi("-121")
	fmt.Println(output)

	output = myAtoi("0")
	fmt.Println(output)
}

func myAtoi(s string) int {

	var output int

	sign := "positive"
	if string(s[0]) == "-" {
		sign = "negtive"
		s = s[1:]
	} else if string(s[0]) == "+" {
		s = s[1:]
	}

	stringLen := len(s)

	for i := 0; i < stringLen; i++ {
		tempNum, _ := strconv.Atoi(string(s[i]))
		output = output*10 + tempNum
	}

	if sign == "negtive" {
		output = output * -1
	}

	return output
}

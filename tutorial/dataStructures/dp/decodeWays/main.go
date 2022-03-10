package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {

	runeArray := []rune(s)
	length := len(runeArray)
	if length == 0 {
		return 0
	}

	if string(runeArray[0]) == "0" {
		return 0
	}

	numwaysArray := make([]int, length)

	numwaysArray[0] = 1

	if length == 1 {
		return numwaysArray[0]
	}

	secondDigit := string(runeArray[0:2])
	num, _ := strconv.Atoi(secondDigit)
	if num > 10 && num <= 19 {
		numwaysArray[1] = 2
	} else if num > 20 && num <= 26 {
		numwaysArray[1] = 2
	} else if num == 10 || num == 20 {
		numwaysArray[1] = 1
	} else if num%10 == 0 {
		numwaysArray[1] = 0
	} else {
		numwaysArray[1] = 1
	}

	for i := 2; i < length; i++ {
		firstDigit := string(runeArray[i])
		if firstDigit != "0" {
			numwaysArray[i] = numwaysArray[i] + numwaysArray[i-1]
		}

		secondDigit := string(runeArray[i-1 : i+1])
		fmt.Println(i)
		fmt.Println(secondDigit)

		num, _ := strconv.Atoi(secondDigit)

		if num >= 10 && num <= 26 {
			numwaysArray[i] = numwaysArray[i] + numwaysArray[i-2]
		}
	}

	return numwaysArray[length-1]

}

func main() {
	output := numDecodings("15")
	fmt.Println(output)
}

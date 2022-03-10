package main

import "fmt"

func lengthOfLastWord(s string) int {
	lenS := len(s)

	lenLastWord := 0
	for i := lenS - 1; i >= 0; {
		for i >= 0 && string(s[i]) == " " {
			i--
		}
		if i < 0 {
			return 0
		}

		for i >= 0 && string(s[i]) != " " {
			//fmt.Println(i)
			//fmt.Println(string(s[i]))
			i--
			lenLastWord++
		}

		return lenLastWord
	}

	return 0
}

func main() {
	length := lengthOfLastWord("computer science")
	fmt.Println(length)

	length = lengthOfLastWord("computer science is a subject")
	fmt.Println(length)

	length = lengthOfLastWord("  ")
	fmt.Println(length)
}

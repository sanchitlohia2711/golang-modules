package main

import "fmt"

func main() {

	var string_first [3]string

	string_first[0] = "abc"
	string_first[1] = "def"
	string_first[2] = "ghi"

	fmt.Println("Output for First Array of string")
	for _, c := range string_first {
		fmt.Println(c)
	}

	string_second := [3]string{
		"ghi",
		"def",
		"abc",
	}

	fmt.Println("\nOutput for Second Array of string")
	for _, c := range string_second {
		fmt.Println(c)
	}
}

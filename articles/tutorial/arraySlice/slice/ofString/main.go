package main

import "fmt"

func main() {

	//First Way
	var string_first []string
	string_first = append(string_first, "abc")
	string_first = append(string_first, "def")
	string_first = append(string_first, "ghi")

	fmt.Println("Output for First slice of string")
	for _, c := range string_first {
		fmt.Println(c)
	}

	//Second Way
	string_second := make([]string, 3)
	string_second[0] = "ghi"
	string_second[1] = "def"
	string_second[2] = "abc"

	fmt.Println("\nOutput for Second slice of string")
	for _, c := range string_second {
		fmt.Println(c)
	}
}

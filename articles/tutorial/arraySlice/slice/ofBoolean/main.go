package main

import "fmt"

func main() {

	//First Way
	var booleans_first []bool
	booleans_first = append(booleans_first, true)
	booleans_first = append(booleans_first, false)
	booleans_first = append(booleans_first, true)

	fmt.Println("Output for First slice of booleans")
	for _, c := range booleans_first {
		fmt.Println(c)
	}

	//Second Way
	booleans_second := make([]bool, 3)
	booleans_second[0] = false
	booleans_second[1] = true
	booleans_second[2] = false

	fmt.Println("\nOutput for Second slice of booleans")
	for _, c := range booleans_second {
		fmt.Println(c)
	}
}

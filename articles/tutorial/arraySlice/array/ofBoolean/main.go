package main

import "fmt"

func main() {

	var booleans_first [3]bool

	booleans_first[0] = true
	booleans_first[1] = false
	booleans_first[2] = true

	fmt.Println("Output for First Array of booleans")
	for _, c := range booleans_first {
		fmt.Println(c)
	}

	booleans_second := [3]bool{
		false,
		true,
		false,
	}

	fmt.Println("\nOutput for Second Array of booleans")
	for _, c := range booleans_second {
		fmt.Println(c)
	}
}

package main

import "fmt"

func main() {

	var integers_first [3]int

	integers_first[0] = 1
	integers_first[1] = 2
	integers_first[2] = 3

	fmt.Println("Output for First Array of integers")
	for _, c := range integers_first {
		fmt.Println(c)
	}

	integers_second := [3]int{
		3,
		2,
		1,
	}

	fmt.Println("\nOutput for Second Array of integers")
	for _, c := range integers_second {
		fmt.Println(c)
	}
}

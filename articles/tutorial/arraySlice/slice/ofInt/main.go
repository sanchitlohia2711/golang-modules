package main

import "fmt"

func main() {

	//First Way
	var integers_first []int
	integers_first = append(integers_first, 1)
	integers_first = append(integers_first, 2)
	integers_first = append(integers_first, 3)

	fmt.Println("Output for First slice of integers")
	for _, c := range integers_first {
		fmt.Println(c)
	}

	//Second Way
	integers_second := make([]int, 3)
	integers_second[0] = 3
	integers_second[1] = 2
	integers_second[2] = 1

	fmt.Println("\nOutput for Second slice of integers")
	for _, c := range integers_second {
		fmt.Println(c)
	}
}

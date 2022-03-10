package main

import "fmt"

func main() {

	var floats_first [3]float64

	floats_first[0] = 1.1
	floats_first[1] = 2.2
	floats_first[2] = 3.3

	fmt.Println("Output for First Array of floats")
	for _, c := range floats_first {
		fmt.Println(c)
	}

	floats_second := [3]float64{
		3.3,
		2.2,
		1.1,
	}

	fmt.Println("\nOutput for Second Array of floats")
	for _, c := range floats_second {
		fmt.Println(c)
	}
}

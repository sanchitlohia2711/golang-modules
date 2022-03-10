package main

import "fmt"

func main() {

	//First Way
	var floats_first []float64
	floats_first = append(floats_first, 1.1)
	floats_first = append(floats_first, 2.2)
	floats_first = append(floats_first, 3.3)

	fmt.Println("Output for First slice of floats")
	for _, c := range floats_first {
		fmt.Println(c)
	}

	//Second Way
	floats_second := make([]float64, 3)
	floats_second[0] = 3.3
	floats_second[1] = 2.2
	floats_second[2] = 1.1

	fmt.Println("\nOutput for Second slice of floats")
	for _, c := range floats_second {
		fmt.Println(c)
	}
}

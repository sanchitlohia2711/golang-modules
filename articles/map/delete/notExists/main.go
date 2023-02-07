package main

import "fmt"

func main() {
	sample := make(map[string]int)
	sample["a"] = 1
	sample["b"] = 2
	sample["c"] = 3

	fmt.Println(sample)

	//Check and delete
	if _, ok := sample["d"]; ok {
		delete(sample, "d")
	}

	fmt.Println(sample)

	//Directly delete
	delete(sample, "d")

	fmt.Println(sample)
}

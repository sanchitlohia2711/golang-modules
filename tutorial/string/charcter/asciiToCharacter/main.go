package main

import "fmt"

func main() {
	sampleASCIIDigits := []int{97, 98, 99}
	for _, digit := range sampleASCIIDigits {
		fmt.Printf("Char %s\n", string(digit))
	}
}

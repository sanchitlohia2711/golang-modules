package main

import "fmt"

func main() {
	var a float64 = 12.0
	var b float32 = a

	fmt.Printf("Underlying Type of b: %T\n", b)

	b2 := float32(a)
	fmt.Printf("Underlying Type of b2: %T\n", b2)
}

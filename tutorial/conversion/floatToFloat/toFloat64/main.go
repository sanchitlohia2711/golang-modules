package main

import "fmt"

func main() {
	var a float32 = 12
	var b float64 = float64(a)

	fmt.Printf("Underlying Type of b: %T\n", b)

	b2 := float64(a)
	fmt.Printf("Underlying Type of b2: %T\n", b2)
}

package main

import "fmt"

func main() {
	var a int = 12
	var b float64 = a

	fmt.Printf("Underlying Type of b: %T\n", b)

	b2 := float64(a)
	fmt.Printf("Underlying Type of b2: %T\n", b2)
}

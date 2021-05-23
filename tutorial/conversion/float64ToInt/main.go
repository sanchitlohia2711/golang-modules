package main

import "fmt"

func main() {
	var a float64 = 12
	var b int = int(a)

	fmt.Printf("Underlying Type of b: %T\n", b)

	b2 := int(a)
	fmt.Printf("Underlying Type of b2: %T\n", b2)
}

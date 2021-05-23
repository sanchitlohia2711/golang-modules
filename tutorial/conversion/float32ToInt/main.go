package main

import "fmt"

func main() {
	var a float32 = 12
	var b int = a

	fmt.Printf("Underlying Type of b: %T\n", b)

	b2 := int(a)
	fmt.Printf("Underlying Type of b2: %T\n", b2)
}

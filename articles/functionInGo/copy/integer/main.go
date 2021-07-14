package main

import "fmt"

func main() {

	src := []int{1, 2, 3, 4, 5}

	dst := make([]int, 5)

	numberOfElementsCopied := copy(dst, src)

	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

	//After changing numbers2
	dst[0] = 10
	fmt.Println("\nAfter changing dst")
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

	//Length of destination is less than length of source
	dst = make([]int, 4)
	numberOfElementsCopied = copy(dst, src)
	fmt.Println("\nLength of dst less than src")
	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

	//Length of destination is greater than length of source
	dst = make([]int, 6)
	numberOfElementsCopied = copy(dst, src)
	fmt.Println("\nLength of dst less than src")
	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

}

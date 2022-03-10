package main

import "fmt"

func pow(x float64, n int) float64 {

	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n == -1 {
		return 1 / x
	}

	val := pow(x, n/2)

	m := x
	if n < 0 {
		m = 1 / x
	}

	if n%2 == 1 || n%2 == -1 {
		return val * val * m
	} else {
		return val * val
	}

}

func main() {
	output := pow(2, 4)
	fmt.Println(output)

	output = pow(2, -4)
	fmt.Println(output)
}

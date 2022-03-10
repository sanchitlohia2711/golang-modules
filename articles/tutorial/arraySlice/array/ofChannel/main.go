package main

import "fmt"

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	var channels_first [3]chan int

	channels_first[0] = channel1
	channels_first[1] = channel2
	channels_first[2] = channel3

	fmt.Println("Output for First Array of channels")
	for _, c := range channels_first {
		fmt.Println(c)
	}

	channel_second := [3]chan int{
		channel1,
		channel2,
		channel3,
	}

	fmt.Println("\nOutput for Second Array of channels")
	for _, c := range channel_second {
		fmt.Println(c)
	}
}

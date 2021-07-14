package main

import "time"

func main() {
	ch := make(chan int)
	test(ch)
	close(ch)
	//ch <- 2
	time.Sleep(time.Second * 5)
}

func test(ch chan int) {
	ch <- 2
}

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go goOne(ch)
	time.Sleep(time.Second * 2)
}

func goOne(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

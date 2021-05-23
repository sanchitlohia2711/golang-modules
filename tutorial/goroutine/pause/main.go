package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	ch := make(chan bool)
	go test(ch)

	time.Sleep(time.Second * 1)
	fmt.Printf("Send Value: %t\n", true)
	ch <- true

	time.Sleep((time.Second * 3))
}

func test(ch chan bool) {
	var value = <-ch
	fmt.Printf("Received Value: %t\n", value)
	fmt.Println("Doing some work")
	utf8.RuneCountInString("as")
}

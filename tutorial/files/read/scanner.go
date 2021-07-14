package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Scanner() {
	file, err := os.Open("../test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

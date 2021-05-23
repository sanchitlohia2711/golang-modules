package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name string
}

func main() {
	j := `{"1":{"Name":"John"}}`
	var b map[int]employee
	json.Unmarshal([]byte(j), &b)

	fmt.Println(b)
}

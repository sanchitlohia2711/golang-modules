package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j := `{"1":"John"}`
	var b map[string]string
	json.Unmarshal([]byte(j), &b)

	fmt.Println(b)
}

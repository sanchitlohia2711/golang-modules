package main

import (
	"fmt"
	"strings"
)

func main() {
	query_param_map := make(map[string]string)

	input := "a=b&x=y"

	input_split := strings.Split(input, "&")

	for _, v := range input_split {
		v_split := strings.Split(v, "=")
		query_param_map[v_split[0]] = v_split[1]

	}
	fmt.Println(query_param_map)

}

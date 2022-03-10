package main

import "fmt"

func main() {
	maps := make([]map[string]string, 3)

	map1 := make(map[string]string)
	map1["1"] = "a"

	map2 := make(map[string]string)
	map2["2"] = "b"

	map3 := make(map[string]string)
	map3["3"] = "c"

	maps[0] = map1
	maps[1] = map2
	maps[2] = map3

	for _, m := range maps {
		fmt.Println(m)
	}
}

package main

import "fmt"

func strStr(haystack string, needle string) int {
	runeHayStack := []rune(haystack)
	runeNeedle := []rune(needle)
	lenHayStack := len(runeHayStack)
	lenNeedle := len(runeNeedle)

	if lenNeedle == 0 && lenHayStack == 0 {
		return 0
	}

	if lenNeedle > lenHayStack {
		return -1
	}

	for i := 0; i <= lenHayStack-lenNeedle; i++ {
		k := i
		j := 0

		for j < lenNeedle {
			if runeHayStack[k] == runeNeedle[j] {
				k = k + 1
				j = j + 1
			} else {
				break
			}
		}

		if j == lenNeedle {
			return i
		}
	}

	return -1

}

func main() {
	output := strStr("lion", "io")
	fmt.Println(output)

	output = strStr("tub", "tus")
	fmt.Println(output)
}

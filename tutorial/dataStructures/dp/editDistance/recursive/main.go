package main

import (
	"fmt"
	"math"
)

func main() {
	output := minDistance("abc", "abcd")
	fmt.Println(output)

	output = minDistance("abc", "ab")
	fmt.Println(output)

	output = minDistance("abc", "abd")
	fmt.Println(output)

	output = minDistance("abce", "abd")
	fmt.Println(output)
}

func minDistance(word1 string, word2 string) int {
	word1Rune := []rune(word1)
	word2Rune := []rune(word2)
	lenWord1 := len(word1Rune)
	lenWord2 := len(word2Rune)

	return minDistanceUtil(word1Rune, word2Rune, lenWord1, lenWord2)
}

func minDistanceUtil(word1 []rune, word2 []rune, lenWord1, lenWord2 int) int {
	if lenWord1 == 0 && lenWord2 == 0 {
		return 0
	}

	if lenWord1 == 0 {
		return lenWord2
	}

	if lenWord2 == 0 {
		return lenWord1
	}

	if word1[lenWord1-1] == word2[lenWord2-1] {
		return minDistanceUtil(word1, word2, lenWord1-1, lenWord2-1)
	} else {
		x := minDistanceUtil(word1, word2, lenWord1-1, lenWord2-1)
		y := minDistanceUtil(word1, word2, lenWord1, lenWord2-1)
		z := minDistanceUtil(word1, word2, lenWord1-1, lenWord2)
		return 1 + minOfThree(x, y, z)
	}
}

func minOfThree(x, y, z int) int {
	output := int(math.Min(float64(x), math.Min(float64(y), float64(z))))
	return output
}

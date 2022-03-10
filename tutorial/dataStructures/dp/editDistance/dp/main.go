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

	editDistanceMatrix := make([][]int, lenWord1+1)

	for i := range editDistanceMatrix {
		editDistanceMatrix[i] = make([]int, lenWord2+1)
	}

	for i := 1; i <= lenWord2; i++ {
		editDistanceMatrix[0][i] = i
	}

	for i := 1; i <= lenWord1; i++ {
		editDistanceMatrix[i][0] = i
	}
	for i := 1; i <= lenWord1; i++ {
		for j := 1; j <= lenWord2; j++ {

			if word1Rune[i-1] == word2Rune[j-1] {
				editDistanceMatrix[i][j] = editDistanceMatrix[i-1][j-1]
			} else {
				editDistanceMatrix[i][j] = 1 + minOfThree(editDistanceMatrix[i-1][j], editDistanceMatrix[i][j-1], editDistanceMatrix[i-1][j-1])
			}
		}
	}
	return editDistanceMatrix[lenWord1][lenWord2]
}

func minOfThree(x, y, z int) int {
	output := int(math.Min(float64(x), math.Min(float64(y), float64(z))))
	return output
}

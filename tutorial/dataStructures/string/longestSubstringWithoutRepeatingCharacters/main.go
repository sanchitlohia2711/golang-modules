package main

import "fmt"

func main() {
	len := lengthOfLongestSubstring("abbabcda")
	fmt.Println(len)
}

func lengthOfLongestSubstring(s string) int {
	charLastIndex := make(map[rune]int)

	longestSubstringLength := 0
	currentSubstringLength := 0
	start := 0

	for index, character := range s {
		lastIndex, ok := charLastIndex[character]
		if !ok || lastIndex < index-currentSubstringLength {
			currentSubstringLength++
		} else {
			if currentSubstringLength > longestSubstringLength {
				longestSubstringLength = currentSubstringLength
			}
			start = lastIndex + 1
			currentSubstringLength = index - start + 1
		}
		charLastIndex[character] = index
	}
	if currentSubstringLength > longestSubstringLength {
		longestSubstringLength = currentSubstringLength
	}
	return longestSubstringLength
}

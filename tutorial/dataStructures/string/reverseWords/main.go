package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseWords(s string) string {

	runeArray := []rune(s)
	length := len(runeArray)

	reverseRuneArray := reverse(runeArray)

	for i := 0; i < length; {
		for i < length && string(reverseRuneArray[i]) == " " {
			i++
		}
		if i == length {
			break
		}
		wordStart := i

		for i < length && string(reverseRuneArray[i]) != " " {
			i++
		}

		wordEnd := i - 1

		reverseRuneArray = reverseIndex(reverseRuneArray, wordStart, wordEnd)

	}

	noSpaceString := strings.TrimSpace(string(reverseRuneArray))
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(noSpaceString, " ")
}

func reverse(s []rune) []rune {
	length := len(s)
	start := 0
	end := length - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	return s
}

func reverseIndex(s []rune, i, j int) []rune {

	start := i
	end := j
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	return s
}

func main() {
	output := reverseWords("hello world")
	fmt.Println(output)

	output = reverseWords("hello")
	fmt.Println(output)
}

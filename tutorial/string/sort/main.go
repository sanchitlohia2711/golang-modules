package main

import (
	"fmt"
	"sort"
)

func main() {
	sortString("bac")
}

func sortString(input string) {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	fmt.Println(string(runeArray))
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

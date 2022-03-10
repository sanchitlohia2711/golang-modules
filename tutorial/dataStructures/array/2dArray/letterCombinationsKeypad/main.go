package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	letterMap := make(map[string][]string)

	letterMap["2"] = []string{"a", "b", "c"}
	letterMap["3"] = []string{"d", "e", "f"}
	letterMap["4"] = []string{"g", "h", "i"}
	letterMap["5"] = []string{"j", "k", "l"}
	letterMap["6"] = []string{"m", "n", "o"}
	letterMap["7"] = []string{"p", "q", "r", "s"}
	letterMap["8"] = []string{"t", "u", "v"}
	letterMap["9"] = []string{"w", "x", "y", "z"}

	runeDigits := []rune(digits)
	length := len(runeDigits)

	temp := ""

	return letterCombinationsUtil(runeDigits, 0, length, temp, letterMap)

}

func letterCombinationsUtil(runeDigits []rune, start, length int, temp string, letterMap map[string][]string) []string {

	if start == length {
		return []string{temp}
	}

	currentDigit := string(runeDigits[start])

	letters := letterMap[currentDigit]

	final := make([]string, 0)
	for _, val := range letters {
		t := temp + val
		output := letterCombinationsUtil(runeDigits, start+1, length, t, letterMap)
		final = append(final, output...)
	}
	return final
}

func main() {

	output := letterCombinations("3")
	fmt.Println(output)

	output = letterCombinations("34")
	fmt.Println(output)

	output = letterCombinations("345")
	fmt.Println(output)
}

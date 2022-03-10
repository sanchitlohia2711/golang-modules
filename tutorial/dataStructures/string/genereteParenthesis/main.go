package main

import "fmt"

func main() {
	output := generateParenthesis(2)
	fmt.Println(output)
}

func generateParenthesis(n int) []string {

	input := ""
	for i := 0; i < n; i++ {
		input = input + "()"
	}
	inputLength := len(input)
	var output []string
	leftParenthesisCount := 0
	generateParenthesisUtil([]rune(input), &output, &leftParenthesisCount, 0, inputLength)
	return output
}

func generateParenthesisUtil(input []rune, output *[]string, leftParenthesisCount *int, start, end int) {

	if start == end {
		if !contains(*output, string(input)) {
			*output = append(*output, string(input))
		}

	} else {
		for i := start; i < end; i++ {
			if input[i] == '(' {
				*leftParenthesisCount++
				temp := input[start]
				input[start] = input[i]
				input[i] = temp
				generateParenthesisUtil(input, output, leftParenthesisCount, start+1, end)
				temp = input[start]
				input[start] = input[i]
				input[i] = temp
				*leftParenthesisCount--
			} else {
				if *leftParenthesisCount > 0 {
					*leftParenthesisCount--
					temp := input[start]
					input[start] = input[i]
					input[i] = temp
					generateParenthesisUtil(input, output, leftParenthesisCount, start+1, end)
					temp = input[start]
					input[start] = input[i]
					input[i] = temp
					*leftParenthesisCount++
				}
			}

		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// func permuteUtil(input []rune, output *[]string, start, end int) {

// 	if start == end {
// 		*output = append(*output, string(input))
// 	} else {
// 		for i := start; i < end; i++ {
// 			temp := input[start]
// 			input[start] = input[i]
// 			input[i] = temp
// 			permuteUtil(input, output, start+1, end)
// 			temp = input[start]
// 			input[start] = input[i]
// 			input[i] = temp

// 		}
// 	}
// }

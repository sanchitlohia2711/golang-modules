package main

import "fmt"

func main() {
	output := isInterleave("aabcc", "dbbca", "aadbbcbcac")
	fmt.Println(output)

	output = isInterleave("", "", "")
	fmt.Println(output)

	output = isInterleave("aaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Println(output)
}
func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Rune := []rune(s1)
	s2Rune := []rune(s2)
	s3Rune := []rune(s3)

	lenS1 := len(s1Rune)
	lenS2 := len(s2Rune)
	lenS3 := len(s3Rune)

	if lenS1+lenS2 != lenS3 {
		return false
	}

	interleavingMatrix := make([][]bool, lenS1+1)

	for i := range interleavingMatrix {
		interleavingMatrix[i] = make([]bool, lenS2+1)
	}

	i := 1
	k := 1

	interleavingMatrix[0][0] = true

	for i <= lenS1 && k <= lenS3 {
		if s1Rune[i-1] == s3Rune[k-1] {
			interleavingMatrix[i][0] = true
			i++
			k++
		} else {
			break
		}
	}

	i = 1
	k = 1

	for i <= lenS2 && k <= lenS3 {
		if s2Rune[i-1] == s3Rune[k-1] {
			interleavingMatrix[0][i] = true
			i++
			k++
		} else {
			break
		}
	}

	for i := 1; i <= lenS1; i++ {
		for j := 1; j <= lenS2; j++ {

			if s1Rune[i-1] == s3Rune[i+j-1] {
				interleavingMatrix[i][j] = interleavingMatrix[i-1][j]
			}

			if s2Rune[j-1] == s3Rune[i+j-1] && !interleavingMatrix[i][j] {
				interleavingMatrix[i][j] = interleavingMatrix[i][j-1]
			}

		}
	}

	return interleavingMatrix[lenS1][lenS2]
}

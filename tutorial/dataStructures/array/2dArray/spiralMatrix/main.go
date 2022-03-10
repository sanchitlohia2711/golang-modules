package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	output := spiralOrder(matrix)
	fmt.Println(output)

	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	output = spiralOrder(matrix)
	fmt.Println(output)
}

func spiralOrder(matrix [][]int) []int {

	var output []int
	totalRows := len(matrix)
	totalColumns := len(matrix[0])

	startRow := 0
	endRow := totalRows - 1
	startColumn := 0
	endColumn := totalColumns - 1

	for startRow <= endRow && startColumn <= endColumn {

		for i := startColumn; i <= endColumn; i++ {
			output = append(output, matrix[startRow][i])
		}

		startRow++

		for i := startRow; i <= endRow; i++ {
			output = append(output, matrix[i][endColumn])
		}

		endColumn--

		if startRow <= endRow {
			for i := endColumn; i >= startColumn; i-- {
				output = append(output, matrix[endRow][i])
			}
		}

		endRow--

		if startColumn <= endColumn {
			for i := endRow; i >= startRow; i-- {
				output = append(output, matrix[i][startColumn])
			}
		}
		startColumn++
	}

	return output

}

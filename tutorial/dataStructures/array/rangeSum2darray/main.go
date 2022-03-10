package main

import "fmt"

type NumMatrix struct {
	matrix     [][]int
	sum_matrix [][]int
	numColumn  int
	numRows    int
}

func initNumArray(matrix [][]int) NumMatrix {
	numRows := len(matrix)
	numColumn := len(matrix[0])

	if numColumn == 0 || numRows == 0 {
		return NumMatrix{}
	}

	sum_matrix := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		sum_matrix[i] = make([]int, numColumn)
	}

	sum_matrix[0][0] = matrix[0][0]
	for i := 1; i < numRows; i++ {
		sum_matrix[i][0] = matrix[i][0] + sum_matrix[i-1][0]
	}

	for i := 1; i < numColumn; i++ {
		sum_matrix[0][i] = matrix[0][i] + sum_matrix[0][i-1]
	}

	for i := 1; i < numRows; i++ {
		for j := 1; j < numColumn; j++ {
			sum_matrix[i][j] = matrix[i][j] + sum_matrix[i-1][j] + sum_matrix[i][j-1] - sum_matrix[i-1][j-1]
		}
	}

	num_matrix := NumMatrix{
		matrix:     matrix,
		sum_matrix: sum_matrix,
		numColumn:  numColumn,
		numRows:    numRows,
	}
	return num_matrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	topSum := 0
	leftSum := 0
	cornerSum := 0
	if row1 > 0 {
		topSum = this.sum_matrix[row1-1][col2]
	}

	if col1 > 0 {
		leftSum = this.sum_matrix[row2][col1-1]
	}

	if row1 > 0 && col1 > 0 {
		cornerSum = this.sum_matrix[row1-1][col1-1]
	}

	return this.sum_matrix[row2][col2] - topSum - leftSum + cornerSum
}

func main() {
	matrix := [][]int{{1, 3, 5}, {6, 7, 4}}
	na := initNumArray(matrix)

	output := na.SumRegion(0, 1, 1, 2)
	fmt.Println(output)

}

package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	paths := make([][]int, len(obstacleGrid))

	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}

	if obstacleGrid[0][0] != 1 {
		paths[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		} else {
			paths[i][0] = paths[i-1][0]
		}

	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		} else {
			paths[0][i] = paths[0][i-1]
		}

	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}

		}
	}

	return paths[m-1][n-1]
}

func main() {
	output := uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	fmt.Println(output)
}

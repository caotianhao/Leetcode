package main

import "fmt"

func minPathSum2099(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 && j != 0 {
				grid[i][j] += grid[0][j-1]
			} else if i != 0 && j == 0 {
				grid[i][j] += grid[i-1][0]
			} else {
				grid[i][j] += min2099(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min2099(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(minPathSum2099([][]int{{0}}))
}

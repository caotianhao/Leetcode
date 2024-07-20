package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	rowMin, colMax, col, min, l, l2 := make([]int, 0), make([]int, 0), make([]int, 0), 0, len(matrix), len(matrix[0])
	for i := 0; i < l; i++ {
		min = 100001
		for j := 0; j < l2; j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
		}
		rowMin = append(rowMin, min)
	}
	//fmt.Println(rowMin)
	//return []int{2}
	for i := 0; i < l; i++ {
		for j := 0; j < l2; j++ {
			if rowMin[i] == matrix[i][j] {
				col = append(col, j)
			}
		}
	}
	//fmt.Println(col)
	for i := 0; i < l; i++ {
		flag := true
		for j := 0; j < l; j++ {
			if matrix[j][col[i]] > rowMin[i] {
				flag = false
			}
		}
		if flag == true {
			colMax = append(colMax, rowMin[i])
		}
	}
	return colMax
}

func main() {
	arr := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	fmt.Println(luckyNumbers(arr))
}

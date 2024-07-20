package main

import "fmt"

func rowAndMaximumOnes(mat [][]int) []int {
	max, index := -1, 0
	for i, v := range mat {
		cnt := 0
		for _, val := range v {
			if val == 1 {
				cnt++
			}
		}
		if cnt > max {
			max, index = cnt, i
		}
	}
	return []int{index, max}
}

func main() {
	fmt.Println(rowAndMaximumOnes([][]int{{0, 1}, {1, 0}}))
}

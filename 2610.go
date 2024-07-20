package main

import "fmt"

func findMatrix(nums []int) [][]int {
	max, map2610 := 0, map[int]int{}
	for _, v := range nums {
		map2610[v]++
		if map2610[v] > max {
			max = map2610[v]
		}
	}
	res := make([][]int, max)
	for i := 0; i < max; i++ {
		tmp := make([]int, 0)
		for k := range map2610 {
			if map2610[k] != 0 {
				tmp = append(tmp, k)
				map2610[k]--
			}
		}
		res[i] = tmp
	}
	return res
}

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
}

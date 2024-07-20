package main

import (
	"fmt"
	"math"
)

func findMaxK(nums []int) int {
	max, map2441 := -1, map[int]int{}
	for _, v := range nums {
		if _, ok := map2441[v]; ok && int(math.Abs(float64(v))) > max {
			max = int(math.Abs(float64(v)))
		}
		map2441[-v] = -1
	}
	return max
}

func main() {
	fmt.Println(findMaxK([]int{-1, 2, -3, 3}))             //3
	fmt.Println(findMaxK([]int{-1, 10, 6, 7, -7, 1}))      //7
	fmt.Println(findMaxK([]int{-10, 8, 6, 7, -2, -3}))     //-1
	fmt.Println(findMaxK([]int{-10, 8, 6, 7, -2, -3, -8})) //8
	fmt.Println(findMaxK([]int{-10}))                      //-1
}

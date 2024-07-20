package main

import "fmt"

func maximumCount(nums []int) int {
	pos, neg := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			neg++
		} else if nums[i] > 0 {
			pos++
		}
	}
	return max2529(pos, neg)
}

func max2529(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumCount([]int{-2, -1, 0, 1}))
}

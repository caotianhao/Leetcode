package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	max, l := 0, len(nums)
	left, right := 0, l-1
	for left < right {
		max = max1877(max, nums[left]+nums[right])
		left++
		right--
	}
	return max
}

func max1877(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minPairSum([]int{3, 5, 4, 2, 4, 6}))
}

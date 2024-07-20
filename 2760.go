package main

import "fmt"

func longestAlternatingSubarray(nums []int, threshold int) int {
	// 暴力枚举所有连续子数组
	max := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i]%2 != 0 {
			continue
		}
	loop:
		for j := i; j < l; j++ {
			part := nums[i : j+1]
			for k := 0; k < len(part)-1; k++ {
				if part[k]%2 == part[k+1]%2 {
					break loop
				}
			}
			for k := 0; k < len(part); k++ {
				if part[k] > threshold {
					break loop
				}
			}
			max = max2760(max, j-i+1)
		}
	}
	return max
}

func max2760(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestAlternatingSubarray([]int{3, 2, 5, 4}, 5))
	fmt.Println(longestAlternatingSubarray([]int{1, 2}, 2))
}
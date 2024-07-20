package main

import "fmt"

func twoSumLessThanK(nums []int, k int) int {
	max := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if t := nums[i] + nums[j]; t < k && t > max {
				max = t
			}
		}
	}
	return max
}

func main() {
	fmt.Println(twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60))
}

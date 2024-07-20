package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	max, sum := 0.0, 0
	for i := range nums {
		if i == k-1 {
			for a := 0; a < k; a++ {
				sum += nums[a]
			}
			max = float64(sum)
		}
		if i >= k {
			sum -= nums[i-k]
			sum += nums[i]
			if float64(sum) > max {
				max = float64(sum)
			}
		}
	}
	return max / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) //12.75
	fmt.Println(findMaxAverage([]int{0, 1, 1, 3, 3}, 4))        //2
}

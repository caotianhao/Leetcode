package main

import "fmt"

func findPrefixScore(nums []int) []int64 {
	l := len(nums)
	max, sum := make([]int, l), make([]int, l)
	max[0] = nums[0]
	for i := 1; i < l; i++ {
		if nums[i] > max[i-1] {
			max[i] = nums[i]
		} else {
			max[i] = max[i-1]
		}
	}
	for i := 0; i < l; i++ {
		sum[i] = max[i] + nums[i]
	}
	res := make([]int64, l+1)
	for i, v := range sum {
		res[i+1] = res[i] + int64(v)
	}
	return res[1:]
}

func main() {
	fmt.Println(findPrefixScore([]int{2, 3, 7, 5, 10}))
}

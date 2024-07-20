package main

import "fmt"

func maximizeSum(nums []int, k int) (sum int) {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	for i := 0; i < k; i++ {
		sum += max
		max++
	}
	return
}

func main() {
	fmt.Println(maximizeSum([]int{1, 2, 3, 6}, 6))
}

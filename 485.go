package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	cnt, max := 0, -1
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}
	}
	if cnt > max {
		max = cnt
	}
	return max
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}

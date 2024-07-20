package main

import "fmt"

func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	} else {
		min, max := 101, 0
		for _, v := range nums {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		for _, v := range nums {
			if v != min && v != max {
				return v
			}
		}
		return -1
	}
}

func main() {
	fmt.Println(findNonMinOrMax([]int{1, 2, 3}))
}

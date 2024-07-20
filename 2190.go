package main

import "fmt"

func mostFrequent(nums []int, key int) int {
	m := map[int]int{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			m[nums[i+1]]++
		}
	}
	max := -1
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	for i, v := range m {
		if v == max {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(mostFrequent([]int{2, 2, 2, 2, 3, 2, 3}, 2))
}

package main

import "fmt"

func mostFrequentEven(nums []int) int {
	hashMap, max, res := map[int]int{}, 0, 100001
	for _, v := range nums {
		if v%2 == 0 {
			hashMap[v]++
		}
	}
	if len(hashMap) == 0 {
		return -1
	}
	for _, v := range hashMap {
		if v > max {
			max = v
		}
	}
	for i, v := range hashMap {
		if v == max {
			res = min2404(res, i)
		}
	}
	return res
}

func min2404(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(mostFrequentEven([]int{0, 1, 2, 2, 4, 4, 1}))
}

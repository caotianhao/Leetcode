package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	min, max := prices[0], -1
	for _, v := range prices {
		if v < min {
			min = v
			continue
		}
		if v-min > max {
			max = v - min
		}
	}
	return max
}

func maxProfit1(prices []int) int {
	n := len(prices)
	dp := make([]int, n)
	min := math.MaxInt64
	max := 0
	for i, p := range prices {
		if p < min {
			min = p
			continue
		}
		dp[i] = p - min
	}
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(maxProfit1([]int{7, 6, 5, 4}))
}

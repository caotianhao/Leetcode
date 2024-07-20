package main

import "fmt"

func maxProduct152(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max152(mx*nums[i], max152(nums[i], mn*nums[i]))
		minF = min152(mn*nums[i], min152(nums[i], mx*nums[i]))
		ans = max152(maxF, ans)
	}
	return ans
}

func max152(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min152(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxProduct152([]int{2, 3, -2, 4}))
}

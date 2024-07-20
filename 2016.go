package main

import "fmt"

func maximumDifference(nums []int) int {
	l, ret := len(nums), -1
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			tmp := nums[j] - nums[i]
			if tmp > 0 {
				ret = max2016(ret, tmp)
			}
		}
	}
	return ret
}

func max2016(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumDifference([]int{7, 1, 5, 4}))  //4
	fmt.Println(maximumDifference([]int{9, 4, 3, 2}))  //-1
	fmt.Println(maximumDifference([]int{1, 5, 2, 10})) //9
}

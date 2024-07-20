package main

import "sort"

func maxProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	return (nums[l-1] - 1) * (nums[l-2] - 1)
}

func main() {

}

package main

import (
	"fmt"
)

type TreeNode654 struct {
	Val   int
	Left  *TreeNode654
	Right *TreeNode654
}

func constructMaximumBinaryTree(nums []int) *TreeNode654 {
	max, ind := findMaxAndIndex(nums)
	if ind != -1 {
		root := &TreeNode654{max, nil, nil}
		root.Left = constructMaximumBinaryTree(nums[:ind])
		root.Right = constructMaximumBinaryTree(nums[ind+1:])
		return root
	}
	return nil
}

func findMaxAndIndex(nums []int) (int, int) {
	max, ind := -1, -1
	for i, v := range nums {
		if v > max {
			max = v
			ind = i
		}
	}
	return max, ind
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
	//fmt.Println(findMaxAndIndex([]int{3, 2, 1, 6, 0, 5}))
}

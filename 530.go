package main

import (
	"fmt"
	"sort"
)

type TreeNode530 struct {
	Val   int
	Left  *TreeNode530
	Right *TreeNode530
}

func getMinimumDifference(root *TreeNode530) int {
	arr, min := make([]int, 0), 999999
	var po = func(node *TreeNode530) {}
	po = func(node *TreeNode530) {
		if node != nil {
			po(node.Left)
			arr = append(arr, node.Val)
			po(node.Right)
		}
	}
	po(root)
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		tmp := arr[i+1] - arr[i]
		if tmp < min {
			min = tmp
			if min == 1 {
				return 1
			}
		}
	}
	return min
}

func main() {
	one := TreeNode530{1, nil, nil}
	three := TreeNode530{3, nil, nil}
	two := TreeNode530{2, &one, &three}
	six := TreeNode530{6, nil, nil}
	four := TreeNode530{4, &two, &six}
	fmt.Println(getMinimumDifference(&four))
}

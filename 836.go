package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	b1 := min836(rec1[2], rec2[2]) > max836(rec1[0], rec2[0])
	b2 := min836(rec1[3], rec2[3]) > max836(rec1[1], rec2[1])
	return b1 && b2
}

func max836(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min836(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
}

package main

import "fmt"

func sumOfDigits(nums []int) int {
	min := 101
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	res := 0
	for min != 0 {
		res += min % 10
		min /= 10
	}
	return (res%2 + 1) % 2
}

func main() {
	fmt.Println(sumOfDigits([]int{34, 23, 1, 24, 75, 33, 54, 8}))
}

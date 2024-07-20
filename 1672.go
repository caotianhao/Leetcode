package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	sum, max := 0, 0
	for i := 0; i < len(accounts); i++ {
		sum = 0
		for j := 0; j < len(accounts[0]); j++ {
			sum += accounts[i][j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	arr := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println(maximumWealth(arr))
}

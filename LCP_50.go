package main

import "fmt"

func giveGem(gem []int, operations [][]int) int {
	min, max := 1001, -1
	for i := 0; i < len(operations); i++ {
		temp := gem[operations[i][0]] / 2
		gem[operations[i][0]] -= temp
		gem[operations[i][1]] += temp
	}
	for i := 0; i < len(gem); i++ {
		if gem[i] > max {
			max = gem[i]
		}
		if gem[i] < min {
			min = gem[i]
		}
	}
	return max - min
}

func main() {
	gem := []int{1250, 700, 400, 8000}
	operations := [][]int{{1, 2}, {3, 1}, {1, 2}}
	fmt.Println(giveGem(gem, operations))
}

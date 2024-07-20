package main

import (
	"fmt"
)

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	totalEn, trainEn, trainEx, tmp, needEx := 0, 0, 0, -1, -10000
	for _, v := range energy {
		totalEn += v
	}
	trainEn = max2383(0, totalEn-initialEnergy+1)
	for _, v := range experience {
		needEx = max2383(needEx, v-tmp)
		tmp += v
	}
	trainEx = max2383(0, needEx-initialExperience)
	return trainEn + trainEx
}

func max2383(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minNumberOfHours(5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1}))
	fmt.Println(minNumberOfHours(2, 4, []int{1}, []int{3}))
}

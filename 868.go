package main

import (
	"fmt"
	"math/bits"
)

func binaryGap(n int) (dis int) {
	if bits.OnesCount(uint(n)) == 1 {
		return 0
	}
	slice868, max := make([]int, 0), 0
	for i, v := range dec2Bin868(n) {
		if v == 1 {
			slice868 = append(slice868, i)
		}
	}
	for i := 0; i < len(slice868)-1; i++ {
		if slice868[i+1]-slice868[i] > max {
			max = slice868[i+1] - slice868[i]
		}
	}
	return max
}

func dec2Bin868(n int) (slice868 []int) {
	for n != 0 {
		slice868 = append(slice868, n%2)
		n /= 2
	}
	return
}

func main() {
	fmt.Println(binaryGap(12))
	//fmt.Println(dec2Bin868(12))
}

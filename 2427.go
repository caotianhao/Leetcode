package main

import "fmt"

func commonFactors(a int, b int) int {
	min, cnt := a, 0
	if b < a {
		min = b
	}
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(commonFactors(12, 6))
}

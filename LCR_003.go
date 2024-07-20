package main

import "fmt"

func countBits2003(n int) []int {
	myArr, cnt := make([]int, 0), 0
	for i := 0; i <= n; i++ {
		cnt = 0
		temp := dec2Bin2003(i)
		for j := 0; j < len(temp); j++ {
			if temp[j] == 1 {
				cnt++
			}
		}
		myArr = append(myArr, cnt)
	}
	return myArr
}

func dec2Bin2003(n int) []int {
	myInt := make([]int, 0)
	if n == 0 {
		return []int{0}
	} else {
		for n != 1 {
			myInt = append(myInt, n%2)
			n /= 2
		}
		myInt = append(myInt, 1)
	}
	return myInt
}

func main() {
	fmt.Println(countBits2003(65))
}
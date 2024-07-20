package main

import "fmt"

func largestAltitude(gain []int) int {
	l, max := len(gain), -999999
	mySlice := make([]int, l+1)
	mySlice[0] = 0
	for i := 0; i < l; i++ {
		mySlice[i+1] = mySlice[i] + gain[i]
	}
	for i := 0; i < len(mySlice); i++ {
		if mySlice[i] > max {
			max = mySlice[i]
		}
	}
	return max
}

func main() {
	arr := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(arr))
}

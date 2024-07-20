package main

import "fmt"

func maxPower(s string) (res int) {
	p, q, l := 0, 0, len(s)
	for q < l {
		if s[q] == s[p] {
			res = max1446(res, q-p+1)
		} else {
			p = q
		}
		q++
	}
	return
}

func max1446(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxPower("leetcode"))
	fmt.Println(maxPower("ll"))
}

package main

import "fmt"

func maxLengthBetweenEqualCharacters(s string) (ans int) {
	l, flag := len(s), true
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				ans = max1624(ans, j-i-1)
				flag = false
			}
		}
	}
	if flag {
		return -1
	}
	return ans
}

func max1624(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("abca"))
	fmt.Println(maxLengthBetweenEqualCharacters("document"))
	fmt.Println(maxLengthBetweenEqualCharacters("a"))
}

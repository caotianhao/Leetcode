package main

import (
	"fmt"
	"strings"
)

func maxScore(s string) int {
	max := -1
	for i := 1; i < len(s); i++ {
		t := strings.Count(s[:i], "0") + strings.Count(s[i:], "1")
		if t > max {
			max = t
		}
	}
	return max
}

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("11"))
}

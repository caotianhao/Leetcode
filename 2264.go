package main

import (
	"fmt"
	"strings"
)

func largestGoodInteger(num string) string {
	flag, max := true, "000"
	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			flag = false
			if strings.Compare(num[i:i+3], max) == 1 {
				max = num[i : i+3]
			}
		}
	}
	if flag {
		return ""
	}
	return max
}

func main() {
	fmt.Println(largestGoodInteger("444666999111555"))
}
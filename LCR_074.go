package main

import (
	"fmt"
	"sort"
)

func merge074(intervals [][]int) (res [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for _, v := range intervals[1:] {
		if v[0] > res[len(res)-1][1] {
			res = append(res, v)
		} else {
			res[len(res)-1][1] = max074(res[len(res)-1][1], v[1])
		}
	}
	return
}

func max074(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(merge074([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

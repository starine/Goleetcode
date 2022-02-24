package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}

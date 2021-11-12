package main

import (
	"fmt"
	"math"
)

/*
45. 猜数字大小 II leetcode375: https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/
*/

func getMoneyAmount(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			f[i][j] = minCost
		}
	}
	return f[1][n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func main() {
	n := 10
	fmt.Println(getMoneyAmount(n))
}

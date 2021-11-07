package main

import "fmt"

/*
范围求和 II leetcode598： https://leetcode-cn.com/problems/range-addition-ii/
*/

func maxCount(m int, n int, ops [][]int) int {
	mina, minb := m, n
	for _, op := range ops {
		mina = min(mina, op[0])
		minb = min(minb, op[1])
	}
	return mina * minb
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	m := 3
	n := 3
	ops := [][]int{{2, 2}, {3, 3}}
	fmt.Println(maxCount(m, n, ops))
}

package main

import (
	"fmt"
	"sort"
)

/*
441. 排列硬币 https://leetcode-cn.com/problems/arranging-coins/
*/
/*从i=1开始，n每次减i，如果不够减，此时i-1就是答案, O（n）
func arrangeCoins(n int) (ans int) {
	for i := 1; n - i >= 0; i++ {
		n -= i
		ans ++
	}
	return
}
*/
//二分查找法
//Search函数采用二分法搜索找到[0, n)区间内最小的满足f(i)==true的值i。
func arrangeCoins(n int) int {
	return sort.Search(n, func(i int) bool {
		i++
		return i*(i+1) > 2*n
	})
}

/*//数学公式 (-b+sqrt(b^2-4ac))/2a
func arrangeCoins(n int) int {
	return int((math.Sqrt(float64(8*n+1)) - 1) / 2)
}
*/
func main() {
	n := 1
	fmt.Println(arrangeCoins(n))
}

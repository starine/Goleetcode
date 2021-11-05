package main

import "fmt"

/*
1218. 最长定差子序列 https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/
*/
func longestSubsequence(arr []int, difference int) (ans int) {
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return
}
func main() {
	arr := []int{1, 5, 7, 8, 5, 3, 4, 2, 1}
	difference := -2
	fmt.Println(longestSubsequence(arr, difference))
}

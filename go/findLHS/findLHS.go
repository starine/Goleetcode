package main

import "fmt"

/*
594. 最长和谐子序列 https://leetcode-cn.com/problems/longest-harmonious-subsequence/
*/

func findLHS(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if cnt[num+1] > 0 && c+cnt[num+1] > ans {
			ans = c + cnt[num+1]
		}
	}
	return
}

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(nums))
}

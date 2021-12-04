package main

import "fmt"

/*
1005. K 次取反后最大化的数组和 https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/
*/

func largestSumAfterKNegations(nums []int, k int) (ans int) {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
		ans += num
	}
	for i := -100; i < 0 && k != 0; i++ {
		if freq[i] > 0 {
			ops := min(k, freq[i])
			ans -= i * ops * 2
			freq[-i] += ops
			k -= ops
		}
	}
	if k > 0 && k%2 == 1 && freq[0] == 0 {
		for i := 1; i <= 100; i++ {
			if freq[i] > 0 {
				ans -= i * 2
				break
			}
		}
	}
	return
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	nums := []int{3, -1, 0, 2}
	k := 3
	fmt.Println(largestSumAfterKNegations(nums, k))
}

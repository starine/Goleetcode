package main

func LIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n) //dp[i]表示以下标i结尾的最长上升子序列长度
	dp[0] = 1
	maxans := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] { //只要前面某个数小于当前数，则要么长度在之前基础上加1，要么保持不变，取较大者
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxans = max(maxans, dp[i])
	}
	return maxans
}

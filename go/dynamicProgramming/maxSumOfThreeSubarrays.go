package main

/*
689. 三个无重叠子数组的最大和 https://leetcode-cn.com/problems/maximum-sum-of-3-non-overlapping-subarrays/
*/

//滑动窗口求单个数组中长度为k的子序列最大的和
func macSumOfOneSubarray(nums []int, k int) (ans []int) {
	var sum1, maxSum1 int
	for i, num := range nums {
		sum1 += num
		if i >= k-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				ans = []int{i - k + 1}
			}
			sum1 -= nums[i-k+1]
		}
	}
	return
}

//滑动窗口求数组中长度为k的两个无重叠子序列最大的和
func macSumOfTwoSubarray(nums []int, k int) (ans []int) {
	var sum1, maxSum1, maxSum1Idx int
	var sum2, maxSum12 int
	for i := k; i < len(nums); i++ {
		sum1 += nums[i-k]
		sum2 += nums[i]
		if i >= k*2-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Idx = i - k*2 + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1 + sum2
				ans = []int{maxSum1Idx, i - k + 1}
			}
			sum1 -= nums[i-k*2+1]
			sum2 -= nums[i-k+1]
		}
	}
	return
}

//滑动窗口求数组中长度为k的三个无重叠子序列最大的和
func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
	var sum1, maxSum1, maxSum1Idx int
	var sum2, maxSum12, maxSum12Idx1, maxSum12Idx2 int
	var sum3, maxTotal int
	for i := k * 2; i < len(nums); i++ {
		sum1 += nums[i-k*2]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= k*3-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Idx = i - k*3 + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1 + sum2
				maxSum12Idx1, maxSum12Idx2 = maxSum1Idx, i-k*2+1
			}
			if maxSum12+sum3 > maxTotal {
				maxTotal = maxSum12 + sum3
				ans = []int{maxSum12Idx1, maxSum12Idx2, i - k + 1}
			}
			sum1 -= nums[i-k*3+1]
			sum2 -= nums[i-k*2+1]
			sum3 -= nums[i-k+1]
		}
	}
	return
}

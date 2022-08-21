/*
* @Author: starine
* @Date:   2022/8/6 23:13
 */

package main

import (
	"fmt"
)

func countBadPairs(nums []int) int64 {
	var cnt int64
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] + 1 == nums[i+1] {
			continue
		}else if nums[i] < nums[i+1] {
			cnt += int64(nums[i+1] - nums[i])
		}else {
			cnt += int64(nums[i] - nums[i+1])
		}
	}
	return cnt
}
//func countBadPairs(nums []int) int64 {
//	dp := make([]int64, len(nums))
//	for i := 1; i < len(nums); i++ {
//		if nums[i] - nums[i-1] != 1 {
//			dp[i] = dp[i-1] + int64(i)
//		}else {
//			dp[i] = dp[i-1]
//		}
//	}
//	return dp[len(nums)-1]
//}
func main() {
	nums := []int{4,1,3,3}
	fmt.Println(countBadPairs(nums))
}

/*
* @Author: starine
* @Date:   2022/8/6 23:48
 */

package main

func minimumReplacement(nums []int) int64 {
	i := len(nums) - 2
	var cnt int64
	right := nums[len(nums) - 1]
	for i >= 0 {
		tmp := nums[i]
		for tmp > right{
			tmp -= right
			cnt++
		}
		right = tmp
		i--
	}
	return cnt
}

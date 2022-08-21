/*
* @Author: starine
* @Date:   2022/8/7 10:34
 */

package main
func arithmeticTriplets(nums []int, diff int) int {
	n := len(nums)
	cnt := 0
	for i := 0; i < n-2; i++ {
		for j := i+1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[j] - nums[i] == diff && nums[k] - nums[j] == diff  {
					cnt++
				}
			}
		}
	}
	return cnt
}

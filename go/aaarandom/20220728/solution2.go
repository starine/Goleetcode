/*
* @Author: starine
* @Date:   2022/8/1 15:40
 */

package main

import (
	"fmt"
	"sort"
)

func longSep(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	max, d := 0,0
	for i, num := range nums {
		for j := 0; j < i; j++ {
			d = num - nums[j]
			if v,ok := dp[j][d]; ok {
				dp[i][d] = v + 1
			}else {
				dp[i][d] = 2
			}
			if max < dp[i][d] {
				max = dp[i][d]
			}
		}
	}
	return  max
}
func main()  {
	nums := []int{9,4,7,1,2,10}
	fmt.Println(longSep(nums))
}
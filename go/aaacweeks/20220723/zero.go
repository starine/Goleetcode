/*
* @Author: starine
* @Date:   2022/7/23 22:59
 */

package main

func zeroFilledSubarray(nums []int) int64 {
	var res, cnt int64
	for _, num := range nums {
		if num == 0 {
			cnt++
		}else {
			res += (1 + cnt) * cnt / 2
			cnt = 0
		}
	}
	if cnt != 0 {
		res += (1 + cnt) * cnt / 2
	}
	return res
}

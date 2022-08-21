package main

import "fmt"

/*
260. 只出现一次的数字 III https://leetcode-cn.com/problems/single-number-iii/
*/

func singleNumber(nums []int) (ans []int) {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	for k, v := range cnt {
		if v == 1 {
			ans = append(ans, k)
		}
	}
	return
}

/*//位运算
func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}
*/

func main() {
	nums := []int{-1, 0}
	fmt.Println(singleNumber(nums))
}

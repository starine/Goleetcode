package main

import "fmt"

/*
求众数 II leetcode229: https://leetcode-cn.com/problems/majority-element-ii/
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
*/

func majorityElement(nums []int) (ans []int) {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for v, c := range cnt {
		if c > len(nums)/3 {
			ans = append(ans, v)
		}
	}
	return
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

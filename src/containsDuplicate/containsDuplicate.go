package main

import "fmt"

/*
217. 存在重复元素 https://leetcode-cn.com/problems/contains-duplicate/
*/
func containsDuplicate(nums []int) bool {
	myset := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := myset[num]; ok {
			return true
		}
		myset[num] = struct{}{}
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

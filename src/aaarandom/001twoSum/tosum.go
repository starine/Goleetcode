package main

import "fmt"

/*
1. 两数之和 ：https://leetcode-cn.com/problems/two-sum/
*/

/*//方法一：两层循环，复杂度为n的平方
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int {i, j}
			}
		}
	}
	return []int{0,0}
}*/

//方法二，利用map
func twoSum(nums []int, target int) []int {
	myMap := map[int]int{}
	for i, num := range nums {
		if j, ok := myMap[target-num]; ok {
			return []int{i, j}
		}
		myMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

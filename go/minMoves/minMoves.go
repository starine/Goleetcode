package main

import "fmt"

/*
453. 最小操作次数使数组元素相等 https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/
*/

/*
每次操作既可以理解为使 n−1 个元素增加 1，也可以理解使 1 个元素减少 1
要计算让数组中所有元素相等的操作数，我们只需要计算将数组中所有元素都减少到数组中元素最小值所需的操作数，
*/
func minMoves(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	ret := 0
	for _, v := range nums {
		ret += v - min
	}
	return ret
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minMoves(nums))
}

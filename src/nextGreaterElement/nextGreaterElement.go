package main

import "fmt"

/*
496. 下一个更大元素 I https://leetcode-cn.com/problems/next-greater-element-i/
*/

func nextGreaterElement(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, m)
	for i, num := range nums1 {
		j := 0
		for j < n && nums2[j] != num {
			j++
		}
		k := j + 1
		for k < n && nums2[k] < nums2[j] {
			k++
		}
		if k < n {
			res[i] = nums2[k]
		} else {
			res[i] = -1
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

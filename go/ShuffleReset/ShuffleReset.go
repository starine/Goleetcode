package main

import (
	"fmt"
	"math/rand"
)

/*
384. 打乱数组 https://leetcode-cn.com/problems/shuffle-an-array/
*/
type Solution struct {
	nums, original []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int(nil), nums...)}
}

func (s *Solution) Reset() []int {
	copy(s.nums, s.original)
	return s.nums
}

func (s *Solution) Shuffle() []int {
	shuffled := make([]int, len(s.nums))
	for i := range shuffled {
		j := rand.Intn(len(s.nums))
		shuffled[i] = s.nums[j]
		s.nums = append(s.nums[:j], s.nums[j+1:]...)
	}
	s.nums = shuffled
	return s.nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	obj := Constructor(nums)
	param_0 := obj.Reset()
	param_1 := obj.Shuffle()
	fmt.Println(param_0)
	fmt.Println(param_1)
}

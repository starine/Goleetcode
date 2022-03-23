package main

import "sort"

func backtrack1(nums []int, used []bool, track []int, res *[][]int) {

	//base case
	if len(track) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i, num := range nums {
		if !used[i] {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			backtrack1(nums, used, append(track, num), res)
			used[i] = false
		}
	}
}
func permuteUnique(num []int) [][]int {
	used := make([]bool, len(num))
	res := [][]int{}
	sort.Ints(num)
	backtrack1(num, used, []int{}, &res)
	return res
}

func main() {

}

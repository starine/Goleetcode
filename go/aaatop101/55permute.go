package main

func backtrack(nums []int, used []bool, track []int, res *[][]int) {

	//base case
	if len(track) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i, num := range nums {
		if !used[i] {
			used[i] = true
			backtrack(nums, used, append(track, num), res)
			used[i] = false
		}
	}
}

func permute(num []int) [][]int {
	used := make([]bool, len(num))
	res := [][]int{}
	backtrack(num, used, []int{}, &res)
	return res
}

func main() {

}

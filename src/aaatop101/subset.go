package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	backtrack4(nums, []int{}, 0, &res)
	return res
}
func backtrack4(nums []int, track []int, start int, res *[][]int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	*res = append(*res, tmp)

	//base case
	if start > len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		backtrack4(nums, append(track, i), i+1, res)
	}
}

func main() {

}

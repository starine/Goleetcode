package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	backtrack3(n, k, []int{}, 1, &res)
	return res
}
func backtrack3(n int, k int, track []int, idx int, res *[][]int) {
	//剪枝:构造的track长度加上区间[idx,n]的长度小于k，往下不可能构造出长度为k的track
	if len(track)+n-idx+1 < k {
		return
	}

	//base case
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	if idx > n {
		return
	}

	for i := idx; i <= n; i++ {
		backtrack3(n, k, append(track, i), i+1, res)
	}
}

func main() {

}

package main

import "sort"

func permutation(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	n := len(t)
	perm := make([]byte, 0, n)
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(i int) {
		//base case
		if i == n {
			ans = append(ans, string(perm))
			return
		}

		for j, b := range vis {
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, t[j])
			backtrack(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return
}

func Permutation(str string) []string {
	res := []string{}

	t := []byte(str)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	used := make([]bool, len(t))
	backtrack5(t, used, []byte{}, &res)

	return res
}
func backtrack5(t []byte, used []bool, track []byte, res *[]string) {
	//base case
	if len(track) == len(t) {
		tmp := make([]byte, len(t))
		copy(tmp, track)
		*res = append(*res, string(tmp))
		return
	}

	for i, b := range t {
		if used[i] || i > 0 && t[i-1] == t[i] && !used[i-1] {
			continue
		}
		used[i] = true
		backtrack5(t, used, append(track, b), res)
		used[i] = false
	}
}

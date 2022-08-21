package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	mp := map[byte]int{}
	n := len(s)
	ans, rp := 0, -1
	for i := 0; i < n; i++ {
		for rp+1 < n && mp[s[rp+1]] == 0 {
			mp[s[rp+1]]++
			rp++
		}
		if rp-i+1 > ans {
			ans = rp - i + 1
		}
		delete(mp, s[i])
	}
	return ans
}
func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

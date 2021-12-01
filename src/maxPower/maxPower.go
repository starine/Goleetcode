package main

import "fmt"

/*
1446. 连续字符 https://leetcode-cn.com/problems/consecutive-characters/
*/

func maxPower(s string) int {
	ans, cnt := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		} else {
			cnt = 1
		}
	}
	return ans
}

func main() {
	s := "abbcccddddeeeeedcba"
	fmt.Println(maxPower(s))
}

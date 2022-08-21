package main

import "fmt"

/*
397. 整数替换 https://leetcode-cn.com/problems/integer-replacement/
*/

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}
	return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	n := 8
	fmt.Println(integerReplacement(n))
}

package main

import "fmt"

/*
318. 最大单词长度乘积 https://leetcode-cn.com/problems/maximum-product-of-word-lengths/
*/

func maxProduct(words []string) (ans int) {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}
	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}

func main() {
	words := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	fmt.Println(maxProduct(words))
}

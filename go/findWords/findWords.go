package main

import (
	"fmt"
	"unicode"
)

/*
500. 键盘行 https://leetcode-cn.com/problems/keyboard-row/
*/

func findWords(words []string) (ans []string) {
	const rowIdx = "12210111011122000010020202"
next:
	for _, word := range words {
		idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
		for _, v := range word[1:] {
			if rowIdx[unicode.ToLower(v)-'a'] != idx {
				continue next
			}
		}
		ans = append(ans, word)
	}
	return
}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}

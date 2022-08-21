package main

import (
	"fmt"
	"unicode"
)

/*
520. 检测大写字母 https://leetcode-cn.com/problems/detect-capital/
*/

func detectCapitalUse(word string) bool {
	if len(word) >= 2 && unicode.IsLower(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
		return false
	}
	for i := 2; i < len(word); i++ {
		if unicode.IsLower(rune(word[i])) != unicode.IsLower(rune(word[1])) {
			return false
		}
	}
	return true
}
func main() {
	word := "USa"
	fmt.Println(detectCapitalUse(word))
}

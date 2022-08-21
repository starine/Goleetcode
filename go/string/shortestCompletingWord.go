package main

/*
748. 最短补全词 https://leetcode-cn.com/problems/shortest-completing-word/
*/
import "unicode"

func shortestCompletingWord(licensePlate string, words []string) (ans string) {
	cnt := [26]int{}
	for _, ch := range licensePlate {
		if unicode.IsLetter(ch) {
			cnt[unicode.ToLower(ch)-'a']++
		}
	}
next:
	for _, word := range words {
		c := [26]int{}
		for _, ch := range word {
			c[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if c[i] < cnt[i] {
				continue next
			}
		}
		if ans == "" || len(word) < len(ans) {
			ans = word
		}
	}
	return
}

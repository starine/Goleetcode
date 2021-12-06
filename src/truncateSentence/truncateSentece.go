package main

import "fmt"

/*
1816. 截断句子 https://leetcode-cn.com/problems/truncate-sentence/
*/

func truncateSentence(s string, k int) string {
	n, end, count := len(s), 0, 0
	for i := 1; i < n; i++ {
		if i == n || s[i] == ' ' {
			count++
			if count == k {
				end = i
				break
			}
		}
	}
	return s[:end]
}
func main() {
	s := "Hello how are you Contestant"
	k := 4
	fmt.Println(truncateSentence(s, k))
}

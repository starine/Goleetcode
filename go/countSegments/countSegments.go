package main

import "fmt"

/*
434. 字符串中的单词数 https://leetcode-cn.com/problems/number-of-segments-in-a-string/
英语中单词都是以空格划分的，例如"Hello, my name is John"中”Hello,“算一个单词
满足单词的第一个下标有以下两个条件：
1.该下标对应的字符不为空格；
2.该下标为初始下标或者该下标的前下标对应的字符为空格；
*/
func countSegments(s string) (ret int) {
	for i, v := range s {
		if (i == 0 || s[i-1] == ' ') && v != ' ' {
			ret++
		}
	}
	return
}

func main() {
	s := "Hello, my name is John"
	fmt.Println(countSegments(s))
}

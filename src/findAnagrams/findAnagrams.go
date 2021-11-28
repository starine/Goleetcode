package main

import "fmt"

/*
438. 找到字符串中所有字母异位词 https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
*/

//超时了^_^
//func findAnagrams(s string, p string) (ans []int) {
//	length := len(p)
//	chp := map[rune]int{}
//	for _, v := range p {
//		chp[v] ++
//	}
//	for i := 0; i < len(s)-length+1;i++{
//		temp := map[rune]int{}
//		for key, value := range chp {
//			temp[key] = value
//		}
//		for j := 0; j < length; j++ {
//			v, ok := temp[rune(s[i+j])]
//			if ok && v > 0{
//				temp[rune(s[i+j])] --
//			}else {
//				break
//			}
//			if j == length-1 {
//				ans  = append(ans, i)
//			}
//		}
//	}
//	return
//}

//在字符串 s 中构造一个长度为与字符串 p 的长度相同的滑动窗口，并在滑动中维护窗口中每种字母的数量；
func findAnagrams(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}

package main

import (
	"fmt"
	"strings"
)

/*
482. 密钥格式化 https://leetcode-cn.com/problems/license-key-formatting/
*/

/*官方解答
func licenseKeyFormatting(s string, k int) string {
	ans := []byte{}
	for i, cnt := len(s)-1, 0; i >= 0; i-- {
		if s[i] != '-' {
			ans = append(ans, byte(unicode.ToUpper(rune(s[i]))))
			cnt++
			if cnt%k == 0 {
				ans = append(ans, '-')
			}
		}
	}
	if len(ans) > 0 && ans[len(ans)-1] == '-' {
		ans = ans[:len(ans)-1]
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}*/

//我的解答 使用strings包提供的函数，更简洁
func licenseKeyFormatting(s string, k int) string {
	ss := strings.ToUpper(strings.ReplaceAll(s, "-", ""))
	n := len(ss)
	var ans []byte
	if n%k != 0 {
		ans = append(append(ans, ss[:n%k]...), '-')
	}
	for i := n % k; i < n; i += k {
		ans = append(append(ans, ss[i:i+k]...), '-')
	}
	if len(ans) > 0 && ans[len(ans)-1] == '-' {
		ans = ans[:len(ans)-1]
	}
	return string(ans)
}

func main() {
	//S := "2-5g-3-J"
	//K := 2
	S := "5F3Z-2e-9-w"
	K := 4
	ret := licenseKeyFormatting(S, K)
	fmt.Println(ret)
}

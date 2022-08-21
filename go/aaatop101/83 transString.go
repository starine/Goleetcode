package main

import "strings"

func trans(s string, n int) string {
	res := strings.Builder{}
	if n == 0 {
		return s
	}
	strs := strings.Split(s, " ")
	i, j := 0, len(strs)-1
	for i < j {
		strs[i], strs[j] = strs[j], strs[i]
		i++
		j--
	}
	for i, str := range strs {
		for _, s := range str {
			b := byte(s)
			res.WriteByte(transByte(b))
		}
		if i != len(strs)-1 {
			res.WriteString(" ")
		}
	}
	result := res.String()
	return result
}

// 大小写转换
func transByte(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - ('a' - 'A')
	} else if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func main() {

}

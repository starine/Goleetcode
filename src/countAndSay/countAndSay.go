package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
38. 外观数列 https://leetcode-cn.com/problems/count-and-say/
要 描述 一个数字字符串，首先要将字符串分割为 最小 数量的组，每个组都由连续的最多 相同字符 组成。
然后对于每个组，先描述字符的数量，然后描述字符，形成一个描述组。
要将描述转换为数字字符串，先将每组中的字符数量用数字替换，再将所有描述组连接起来。
*/
func countAndSay(n int) string {
	var say string = "1"

	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(say); start = j {
			for j < len(say) && say[start] == say[j] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(say[start])
		}
		say = cur.String()
	}
	return say
}

func main() {
	input := 4
	fmt.Println(countAndSay(input))
}

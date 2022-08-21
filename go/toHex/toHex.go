package main

import (
	"fmt"
	"strings"
)

/*
位运算
405. 数字转换为十六进制数:https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/
题目要求将给定的整数num 转换为十六进制数，负整数使用补码运算方法。
符号位是 0 表示正整数和零，符号位是 1 表示负整数。
32 位有符号整数的二进制数有 32 位，一位十六进制数对应四位二进制数，因此 32 位有符号整数的十六进制数有 8 位。
假设二进制数的 8 组从低位到高位依次是第 0 组到第 7 组，则对于第 i 组，可以通过num >> (4 * i) & 0xf得到该组的值，

其取值范围是 0 到 15。将每一组的值转换为十六进制数的做法如下：
对于 0 到 9，数字本身就是十六进制数；
对于 10 到 15，将其转换为 a 到 f 中的对应字母。

对于负整数，由于最高位一定不是 0，因此不会出现前导零。对于零和正整数，可能出现前导零。避免前导零的做法如下：
如果 num=0，则直接返回 0；
如果num>0，则在遍历每一组的值时，从第一个不是 0 的值开始拼接成十六进制数。
*/

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	sb := &strings.Builder{}
	for i := 7; i >= 0; i-- {
		val := num >> (4 * i) & 0xf
		if val > 0 || sb.Len() > 0 {
			var digit byte
			if val < 10 {
				digit = '0' + byte(val)
			} else {
				digit = 'a' + byte(val-10)
			}
			sb.WriteByte(digit)
		}

	}
	return sb.String()
}

func main() {
	var input int = -26
	ret := toHex(input)
	fmt.Println(ret)
}

package main

import (
	"fmt"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
166. 分数到小数 https://leetcode-cn.com/problems/fraction-to-recurring-decimal/
*/
/*
长除法
1.如果结果是负数则将负号拼接到结果中，如果结果是正数则跳过这一步；
2.将整数部分拼接到结果中；
3.将小数点拼接到结果中
4.计算小数部分时，每次将余数乘以 10，然后计算小数的下一位数字，并得到新的余数。重复上述操作直到余数变成 0 或者找到循环节。
如果余数变成 0，则结果是有限小数，将小数部分拼接到结果中。
如果找到循环节，则找到循环节的开始位置和结束位置并加上括号，然后将小数部分拼接到结果中。
5.循环节：对于相同的余数，计算得到的小数的下一位数字一定是相同的，因此如果计算过程中发现某一位的余数在之前已经出现过，则为找到循环节。
为了记录每个余数是否已经出现过，需要使用哈希表存储每个余数在小数部分第一次出现的下标。
*/
func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	s := []byte{}
	if numerator < 0 != (denominator < 0) {
		s = append(s, '-')
	}

	// 整数部分
	numerator = abs(numerator)
	denominator = abs(denominator)
	integerPart := numerator / denominator
	s = append(s, strconv.Itoa(integerPart)...)
	s = append(s, '.')

	// 小数部分
	indexMap := map[int]int{}
	remainder := numerator % denominator
	for remainder != 0 && indexMap[remainder] == 0 {
		indexMap[remainder] = len(s)
		remainder *= 10
		s = append(s, '0'+byte(remainder/denominator))
		remainder %= denominator
	}
	if remainder > 0 { // 有循环节
		insertIndex := indexMap[remainder]
		s = append(s[:insertIndex], append([]byte{'('}, s[insertIndex:]...)...)
		s = append(s, ')')
	}

	return string(s)
}

func main() {
	var numerator int = 1
	var denominator int = 3
	ret := fractionToDecimal(numerator, denominator)
	fmt.Println(ret)
}

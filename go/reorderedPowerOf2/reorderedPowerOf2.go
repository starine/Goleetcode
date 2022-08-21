package main

import "fmt"

/*
869. 重新排序得到 2 的幂 https://leetcode-cn.com/problems/reordered-power-of-2/
*/

func countDigits(n int) (cnt [10]int) {
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}
	return
}

var powerOf2Digits = map[[10]int]bool{}

func init() {
	for n := 1; n < 1e9; n <<= 1 {
		powerOf2Digits[countDigits(n)] = true
	}
}
func reorderedPowerOf2(n int) bool {
	return powerOf2Digits[countDigits(n)]
}

func main() {
	N := 46
	fmt.Println(reorderedPowerOf2(N))
}

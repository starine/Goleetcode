package main

import "fmt"

/*
66. 加一 https://leetcode-cn.com/problems/plus-one/
*/

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	//digits中所有数字都为9
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

func main() {
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}

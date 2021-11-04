package main

import (
	"fmt"
)

/*
367. 有效的完全平方数 https://leetcode-cn.com/problems/valid-perfect-square/
*/

/*//方法一：使用sqrt
func isPerfectSquare(num int) bool {
	x := int(math.Sqrt(float64(num)))
	return x*x == num
}*/

/*//方法二
func isPerfectSquare(num int) bool {
	for x := 1; x*x <= num; x++ {
		if x*x == num {
			return true
		}
	}
	return false
}*/
//方法三，二分查找
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square < num {
			left = mid + 1
		} else if square > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	num := 16
	fmt.Println(isPerfectSquare(num))
}

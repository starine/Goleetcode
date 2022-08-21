package main

import "fmt"

/*
372. 超级次方 https://leetcode-cn.com/problems/super-pow/
*/
const mod = 1337

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func superPow(a int, b []int) int {
	ans := 1
	for i := len(b) - 1; i >= 0; i-- {
		ans = ans * pow(a, b[i]) % mod
		a = pow(a, 10)
	}
	return ans
}

func main() {
	a := 2147483647
	b := []int{2, 0, 0}
	fmt.Println(superPow(a, b))
}

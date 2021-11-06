package main

import "fmt"

/*
丢失的数字 leetcode268：https://leetcode-cn.com/problems/missing-number/
*/
/*func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrsum := 0
	for _, num := range nums {
		arrsum += num
	}
	return total - arrsum
}*/
func missingNumber(nums []int) (xor int) {
	for i, num := range nums {
		xor ^= i ^ num
	}
	return xor ^ len(nums)
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

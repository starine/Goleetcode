package main

import (
	"fmt"
	"strconv"
)

//func jumpFloor( number int ) int {
//	res := 0
//	jump(number,0,&res)
//	return res
//}
//func jump(n int, track int, res *int)  {
//	if track >= n {
//		*res++
//		return
//	}
//	for i := 1; i <=2; i++ {
//		track += i
//		jump(n,track,res)
//	}
//}
//超内存了

func jumpFloor(number int) int {
	if number < 2 {
		return number
	}
	p, q, r := 1, 1, 1
	for i := 2; i <= number; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	var s1 string = "0000"
	num1, _ := strconv.ParseInt(s1, 16, 16)
	fmt.Println(num1)
}

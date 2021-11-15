package main

import (
	"fmt"
	"math"
)

/*
48. 灯泡开关 leetcode319: https://leetcode-cn.com/problems/bulb-switcher/
*/

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}
func main() {
	n := 9
	fmt.Println(bulbSwitch(n))
}

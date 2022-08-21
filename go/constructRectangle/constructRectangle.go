package main

import (
	"fmt"
	"math"
)

/*
构造矩形 leetcode492: https://leetcode-cn.com/problems/construct-the-rectangle/
L >= W
L 和 W 之间的差距应当尽可能小
*/

func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w > 0 {
		w--
	}
	return []int{area / w, w}
}

func main() {
	area := 4
	fmt.Println(constructRectangle(area))
}

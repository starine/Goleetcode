package main

import (
	"fmt"
)
/*
题目描述：https://leetcode-cn.com/problems/rectangle-area/
两个矩形覆盖的总面积等于两个矩形的面积之和减去两个矩形的重叠部分的面积。由于两个矩形的左下顶点和右上顶点已知，因此两个矩形的面积可以直接计算。
如果两个矩形重叠，则重叠部分的水平边投影到 x 轴上的区间为 [ max(ax1,bx1), min(ax2,bx2) ]
y轴同理
只有当两条线段的长度都大于 0 时，重叠部分的面积才大于 0，否则重叠部分的面积为 0。
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	ret := (ax2 - ax1) * (ay2 - ay1) + (bx2 - bx1) * (by2 - by1)
	overlapWidth := min(ax2, bx2) - max(ax1, bx1)
	overlapHeight := min(ay2, by2) - max(ay1, by1)
	return ret - max(overlapWidth, 0) * max(overlapHeight, 0)
}
func main()  {
	ret := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)
	fmt.Println(ret)
}

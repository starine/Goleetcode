/*
* @Author: starine
* @Date:   2022/8/19 15:20
 */

package main

import (
	"fmt"
	"sort"
)

//实现一个方法，输入2个参数，返回一个数值。
//int executeTask(list<int> listTask, threadCount n)
//第一个参数是一个list<int>, 第二个参数是一个int类型的n。
//意思是启动n个线程来执行list中的任务，每个任务用时就是其本身值相等的秒数。请计算通过n个线程执行完这个list需要多少秒。
//
//示例：
//输入：list<1,2,3>, 3
//返回：3
//输入：list<1,2,3,4>, 3
//返回：5
//输入：list<1,2,3,5,2,2,6,7,8>  3
//返回：14
//输入：list<1,2,5,5,2,2,6,7,8>  4
//返回：13

//实现一个方法，输入一个数组, 数组中每个元素值为一对begin和end。每个begin和end表示申请一段内存的起始，begin和end间可能重叠，请返回一个数值，表示共申请了多少内存，重叠部分不累计。
//例：
//input： [2,4], [4,6], [7,17], [10,15]
//output:14
//
//input： [2,4], [4,6], [14,17], [10,15]
//output: 11

func merge(list [][]int) int{
	count := 0
	res := [][]int{}
	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0]
	})
	res = append(res, list[0])
	for i := 0; i < len(list); i++ {
		if list[i][0] < res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], list[i][1])
		}else {
			res = append(res, list[i])
		}
	}
	for i := 0; i < len(res); i++ {
		tmp := res[i][1] - res[i][0]
		count += tmp
	}
	return count
}
func max(a, b int) int  {
	if a >= b {
		return a
	}
	return b
}
func main()  {
	input := [][]int{{2,4}, {4,6}, {7,17}, {10,15}}
	fmt.Println(merge(input))

	input2 := [][]int{{2,4}, {4,6}, {14,17}, {10,15}}
	fmt.Println(merge(input2))
}
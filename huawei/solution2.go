/*
* @Author: starine
* @Date:   2022/8/19 16:52
 */

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

package main

import (
	"fmt"
	"math"
)

func executeTask(listTask []int, n int) int {
	ans := 0
	if len(listTask) <= n {
		return maxOfNums(listTask)
	}
	thread := make([]int, n)
	r := 0 //r表示此时从list中取出的线程的索引
	for ; r < n; r++ {
		thread[r] = listTask[r]
	}
	for r < len(listTask) {
		min := minOfNums(thread)
		ans += min //当线程中需要时间片最小的任务执行完成，会加入新的线程
		for i := 0; i < n; i++ {
			thread[i] = thread[i] - min
			if thread[i] == 0 && r < len(listTask){
				thread[i] = listTask[r]
				r++
			}
		}
	}
	ans += maxOfNums(thread)
	return ans
}
func minOfNums(nums []int) int {
	min := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
func maxOfNums(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
func main()  {
	list1 := []int{1,2,3}
	list2 := []int{1,2,3,4}
	list3 := []int{1,2,3,5,2,2,6,7,8}
	list4 := []int{1,2,5,5,2,2,6,7,8}
	fmt.Println(executeTask(list1, 3))
	fmt.Println(executeTask(list2, 3))
	fmt.Println(executeTask(list3, 3))
	fmt.Println(executeTask(list4, 4))
}



/*
* @Author: starine
* @Date:   2022/6/24 23:26
 */

/*
给定一个只包含0～9的无重复数字的数组arr；

给定一个数字limit；

要求用数组arr中的数字构造一个比limit小的最大整数，数组arr中的数字可以重复使用。

*/

package main

import (
	"fmt"
	"sort"
)

func maxLimit(arr []int, limit int) int {
	//先排序，方便通过索引找离某个数最近的数字
	sort.Ints(arr)
	limit--

	//小技巧，用offset来从limit中取单个数字
	offset := 1
	for offset <= limit/10 {
		offset *= 10
	} //到这一步，offset和limit同位数，例如limit=221，则offset=100

	//定义一个递归函数，从`limit-1`中从左往右每次取一个数字`cur`，去构造最接近他的数
	//如果返回的数做不到和limit位数一样，返回-1
	//如果能做到，返回的结果就是答案
	ans := recursive(arr, limit, offset)

	if ans != -1 { //已经构造出一个符合要求的、和limit位数一样的数字
		return ans
	} else { //ans = -1，构造一个比limit小一位的最大数字
		return rest(arr, offset/10)
	}
}

func recursive(arr []int, limit, offset int) int {
	//构造出的数字完全等于limit，即一直是走分支 arr[near] == cur
	if offset == 0 {
		return limit
	}

	cur := (limit / offset) % 10

	//从数组中找一个最接近cur的数字near,返回数组索引
	near := findNear(arr, cur)

	if near == -1 { //如果数组中找不到，向上层返回-1
		return -1
	} else if arr[near] == cur { //数组中找到数字和当前数字一样，递归向下
		ans := recursive(arr, limit, offset/10)
		if ans != -1 {
			return ans
		} else if near > 0 { //后面的位数无法构造出小于limit的，但是我可以让这一位的数字更小 < cur,后面的数字全选最大的
			near--
			return (limit/(offset*10))*offset*10 + (arr[near] * offset) + rest(arr, offset/10)
		} else { //后面的位数无法构造出小于limit的，而且这一位找不到更小的了，返回-1，去上一层找
			return -1
		}
	} else { //arr[near] < cur
		return (limit/(offset*10))*offset*10 + (arr[near] * offset) + rest(arr, offset/10)
	}
}

//比如offset = 10000，5位数字，把arr中最大的数字x，拼成xxxxx，返回
func rest(arr []int, offset int) int {
	res := 0
	for offset > 0 {
		res += arr[len(arr)-1] * offset
		offset /= 10
	}
	return res
}

// 在有序数组arr中，找到<=num，且最大的数字，返回该数字的索引
// 如果所有数字都大于num，返回-1
func findNear(arr []int, num int) int {
	l, r := 0, len(arr)-1
	m, near := 0, -1
	for l <= r {
		m = l + (r-l)>>1
		if arr[m] <= num {
			near = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return near
}

//测试
func main() {
	arr := []int{2, 6, 7}
	limit := 261
	fmt.Println(maxLimit(arr, limit))

	limit2 := 222
	fmt.Println(maxLimit(arr, limit2))
}

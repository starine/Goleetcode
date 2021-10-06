package main

import (
	"fmt"
)

/*
414. 第三大的数 https://leetcode-cn.com/problems/third-maximum-number/
方法一：排序
将数组从大到小排序后，从头开始遍历数组，通过判断相邻元素是否不同，来统计不同元素的个数。
如果能找到三个不同的元素，就返回第三大的元素，否则返回最大的元素。
时间复杂度：O(nlogn)，其中 n 是数组nums 的长度。排序需要O(nlogn) 的时间。
空间复杂度：O(logn)。排序需要的栈空间为 O(logn)。

func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, diff := 1, 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			diff++
			if diff == 3 { // 此时 nums[i] 就是第三大的数
				return nums[i]
			}
		}
	}
	return nums[0]
}
*/

/*
方法二：有序集合
我们可以遍历数组，同时用一个有序集合来维护数组中前三大的数。具体做法是每遍历一个数，就将其插入有序集合，
若有序集合的大小超过 3，就删除集合中的最小元素。这样可以保证有序集合的大小至多为 3，且遍历结束后，若有序集合的大小为 3，
其最小值就是数组中第三大的数；若有序集合的大小不足 3，那么就返回有序集合中的最大值。
时间复杂度：O(n)，其中 n 是数组 nums 的长度。由于有序集合的大小至多为 3，插入和删除的时间复杂度可以视作是O(1) 的，因此时间复杂度为 O(n)。
空间复杂度：O(1)。

func thirdMax(nums []int) int {
	t := redblacktree.NewWithIntComparator()
	for _, num := range nums {
		t.Put(num, nil)
		if t.Size() > 3 {
			t.Remove(t.Left().Key)
		}
	}
	if t.Size() == 3 {
		return t.Left().Key.(int)
	}
	return t.Right().Key.(int)
}
*/
/*
方法三：一次遍历
我们可以遍历数组，并用三个变量 a、b 和 c 来维护数组中的最大值、次大值和第三大值，以模拟方法二中的插入和删除操作。
为方便编程实现，我们将其均初始化为小于数组最小值的元素，视作「无穷小」，比如 -2^63等。
遍历数组，对于数组中的元素 num：
若 num>a，我们将 c 替换为 b，b 替换为 a，a 替换为 num，这模拟了将 num 插入有序集合，并删除有序集合中的最小值的过程；
若 a>num>b，类似地，我们将 c 替换为 b，b 替换为 num，a 保持不变；
若 b>num>c，类似地，我们将 c 替换为 num，a 和 b 保持不变；
其余情况不做处理。
遍历结束后，若 cc 仍然为 -2^63，则说明数组中不存在三个或三个以上的不同元素，即第三大的数不存在，返回 a，否则返回 c。

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if a > num && num > b {
			b, c = num, b
		} else if b > num && num > c {
			c = num
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
*/
/*
另一种不依赖元素范围的做法是，将 a、b 和 c 初始化为空指针或空对象，视作「无穷小」，并在比较大小前先判断是否为空指针或空对象。
遍历结束后，若 c 为空，则说明第三大的数不存在，返回 a，否则返回 c。
*/
func thirdMax(nums []int) int {
	var a, b, c *int
	for _, num := range nums {
		num := num
		if a == nil || num > *a {
			a, b, c = &num, a, b
		} else if *a > num && (b == nil || num > *b) {
			b, c = &num, b
		} else if b != nil && *b > num && (c == nil || num > *c) {
			c = &num
		}
	}
	if c == nil {
		return *a
	}
	return *c
}

func main() {
	nums := []int{2, 2, 3, 1}
	fmt.Println(thirdMax(nums))
}

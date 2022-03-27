package main

func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		mid := low + (high-low)>>1
		if numbers[mid] < numbers[high] { //最⼩的数字在mid左边
			high = mid
		} else if numbers[mid] > numbers[high] { //最⼩的数字在mid右边
			low = mid + 1
		} else { //⽆法判断，⼀个⼀个试
			high--
		}
	}
	return numbers[low]
}

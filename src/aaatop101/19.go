package main

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right { //⼆分法
		mid := left + (right-left)>>1
		if nums[mid] > nums[mid+1] { //右边是往下，不⼀定有坡峰
			right = mid
		} else { //右边是往上，⼀定能找到波峰
			left = mid + 1
		}
	}
	return right //其中⼀个波峰
}

func main() {

}

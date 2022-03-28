package main

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1

	for {
		pivot := partition(nums, left, right)
		if pivot < len(nums)-k {
			left = pivot + 1
		} else if pivot > len(nums)-k {
			right = pivot - 1
		} else {
			return nums[pivot]
		}
	}
}

// 选取基准，将待排序序列拆分成较小、较大的两个子序列
func partition(nums []int, left, right int) int {
	pivot := left
	i, j := pivot+1, pivot+1
	for j <= right {
		if nums[j] <= nums[pivot] {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
		j++
	}
	nums[pivot], nums[i-1] = nums[i-1], nums[pivot]
	return i - 1
}

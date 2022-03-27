package main

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

//归并排序，合并过程统计逆序对
func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)>>1
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	tmp := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] { //nums[i]大于nums[j]之前的数，即逆序
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1 //左边剩下的数全部大于右边的数
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}

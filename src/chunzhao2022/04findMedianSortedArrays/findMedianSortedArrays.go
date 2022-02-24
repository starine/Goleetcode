package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length&1 == 1 {
		midIndex := length / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := length/2-1, length/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2
	}
	return 0
}
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 < pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	nums := make([]int, m+n)
	index := 0
	for i < m && j < n {
		if nums1[i] >= nums2[j] {
			nums[index] = nums2[j]
			index++
			j++
		} else if nums1[i] < nums2[j] {
			nums[index] = nums1[i]
			index++
			i++
		}
	}
	for i < m {
		nums[index] = nums1[i]
		index++
		i++
	}
	for j < n {
		nums[index] = nums2[j]
		index++
		j++
	}
	length := m + n
	if length&1 == 1 {
		return float64(nums[length/2])
	} else {
		return float64(nums[length/2-1]+nums[length/2]) / 2
	}
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	index1, index2 := 0, 0
	left, right := -1, -1
	for i := 0; i <= length/2; i++ {
		left = right
		if index1 < m && (index2 >= n || nums1[index1] < nums2[index2]) {
			right = nums1[index1]
			index1++
		} else {
			right = nums2[index2]
			index2++
		}
	}
	if length&1 == 1 {
		return float64(right)
	} else {
		return float64(left+right) / 2
	}

}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays3(nums1, nums2))
	fmt.Println(13 / 2)
}

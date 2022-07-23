/*
* @Author: starine
* @Date:   2022/7/19 11:46
 */

package main

import "fmt"

func quickSort(arr []int) []int {
	return quick(arr, 0, len(arr)-1)
}
func quick(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		quick(arr, left, partitionIndex-1)
		quick(arr, partitionIndex+1, right)
	}
	return arr
}
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

func main() {
	arr := []int{63, 43, 3, 2, 67, 89, 100}
	fmt.Println(quickSort(arr))
}

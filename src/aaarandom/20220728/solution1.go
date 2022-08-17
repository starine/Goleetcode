/*
* @Author: starine
* @Date:   2022/7/28 14:27
 */

package main

import (
	"fmt"
)

//func minimumOperations(nums []int) int {
//	sort.Ints(nums)
//	slow := 1
//	for fast := 1;fast < len(nums);fast++{
//		if nums[fast] != nums[fast-1] {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	if nums[0] == 0 {
//		return slow - 1
//	}
//	return slow
//}

func main()  {
	nums := []int{5,3,1,0,2,4,5}
	fmt.Println(closestMeetingNode(nums,3,2))
}
func closestMeetingNode(edges []int, node1 int, node2 int) int {
	res := len(edges) + 1
	next := make(map[int]int)
	for k,v := range edges {
		next[k] = v
	}
	vis1 := make(map[int]bool)
	vis2 := make(map[int]bool)
	i := node1
	for !vis1[i] && i != node2 && i != -1 {
		vis1[i] = true
		i = next[i]
	}
	if i == node2 {
		vis1[node2] = true
	}
	j := node2
	for !vis2[j] && j != node1 && j != -1 {
		vis2[j] = true
		j = next[j]
	}
	if j == node1 {
		vis2[node1] = true
	}
	for k, _ := range vis1 {
		if vis2[k] && k < res{
			res = k
		}
	}
	if res == len(edges) + 1 {
		return -1
	}
	return res
}
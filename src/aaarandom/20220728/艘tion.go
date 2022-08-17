/*
* @Author: starine
* @Date:   2022/8/6 22:34
 */

package main

import (
	"fmt"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var res [][]int
	sort.Slice(items1, func(i, j int) bool {
		return items1[i][0] < items1[j][0]
	})
	sort.Slice(items2, func(i, j int) bool {
		return items2[i][0] < items2[j][0]
	})
	fmt.Println(items1)
	fmt.Println(items2)
	fmt.Println(len(items1), len(items2))
	i, j := 0, 0
	for i < len(items1) && j < len(items2) {
		if items1[i][0] == items2[j][0] {
			res = append(res,[]int{items1[i][0],items1[i][1]+items2[j][1]})
			i++
			j++
		}else if items1[i][0] < items2[j][0] {
			res = append(res,[]int{items1[i][0],items1[i][1]})
			i++
		}else {
			res = append(res,[]int{items2[j][0],items2[j][1]})
			j++
		}
	}
	fmt.Println(i, j)
	fmt.Println(len(items1), len(items2))
	for i < len(items1) {
		res = append(res,[]int{items1[i][0],items1[i][1]})
		i++
	}
	for j < len(items2) {
		res = append(res,[]int{items2[j][0],items2[j][1]})
		j++
	}
	fmt.Println(i, j)
	return res
}
func main()  {
	items1 := [][]int{{5,1},{4,2},{3,3},{2,4},{1,5}}
	items2 :=[][]int{{7,1},{6,2},{5,3},{4,4}}
	fmt.Println(mergeSimilarItems(items1,items2))
}
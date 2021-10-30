package main

import "fmt"

/*
路径交叉 leetcode335: https://leetcode-cn.com/problems/self-crossing/
*/
func isSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		if i == 4 && distance[3] == distance[1] && distance[4] >= distance[2]-distance[0] {
			return true
		}
		if i >= 5 && distance[i-3]-distance[i-5] <= distance[i-1] && distance[i-1] <= distance[i-3] &&
			distance[i-2]-distance[i-4] <= distance[i] && distance[i-2] > distance[i-4] {
			return true
		}
	}
	return false
}

func main() {
	distance := []int{2, 1, 1, 2}
	fmt.Println(isSelfCrossing(distance))
}

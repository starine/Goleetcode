package main

import "fmt"

/*
575. 分糖果 https://leetcode-cn.com/problems/distribute-candies/
*/

func distributeCandies(candyType []int) int {
	set := map[int]struct{}{}
	for _, v := range candyType {
		set[v] = struct{}{}
	}
	ans := len(candyType) / 2
	if len(set) < ans {
		ans = len(set)
	}
	return ans
}

func main() {
	candies := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies(candies))
}

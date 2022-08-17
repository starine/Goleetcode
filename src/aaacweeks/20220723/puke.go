/*
* @Author: starine
* @Date:   2022/7/23 22:34
 */

package main

import "fmt"

func bestHand(ranks []int, suits []byte) string {
	if suits[0] == suits[1] && suits[0]== suits[2] && suits[0]== suits[3] && suits[0]== suits[4] {
		return "Flush"
	}
	cnt := make(map[int]int)
	for _, rank := range ranks {
		cnt[rank]++
	}
	max := 0
	for _, c := range cnt {
		if c > max {
			max = c
		}
	}
	if max >= 3 {
		return "Three of a Kind"
	}
	if max == 2 {
		return "Pair"
	}
	return "High Card"
}

func main() {
	ranks := []int{4,4,2,4,4}
	suits := []byte{'d','a','a','b','c'}
	fmt.Println(bestHand(ranks,suits))
}
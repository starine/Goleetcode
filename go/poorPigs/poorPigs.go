package main

import (
	"fmt"
	"math"
)

/*
458. 可怜的小猪 https://leetcode-cn.com/problems/poor-pigs/
*/

func poorPigs(buckets, minutesToDie, minutesToTest int) int {
	states := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(states))))
}

func main() {
	buckets := 1000
	minutesToDie := 15
	minutesToTest := 60
	fmt.Println(poorPigs(buckets, minutesToDie, minutesToTest))
}

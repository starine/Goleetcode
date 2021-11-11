package main

import "fmt"

/*
43. 提莫攻击 leetcode495:  https://leetcode-cn.com/problems/teemo-attacking/
*/
func findPoisonedDuration(timeSeries []int, duration int) (ans int) {
	expired := 0
	for _, t := range timeSeries {
		if t >= expired {
			ans += duration
		} else {
			ans += t + duration - expired
		}
		expired = t + duration
	}
	return
}
func main() {
	timeSeries := []int{1, 2}
	duration := 2
	fmt.Println(findPoisonedDuration(timeSeries, duration))
}

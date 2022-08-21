package main

func MoreThanHalfNum_Solution(numbers []int) int {
	cnt := map[int]int{}
	h := len(numbers) >> 1
	for _, num := range numbers {
		cnt[num]++
		if cnt[num] > h {
			return num
		}
	}
	return -1
}

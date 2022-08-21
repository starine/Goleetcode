package main

import "sort"

func singleNumber(nums []int) (ans []int) {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	for num, occ := range freq {
		if occ == 1 {
			ans = append(ans, num)
		}
	}
	sort.Ints(ans)
	return
}

func singleNumber2(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}

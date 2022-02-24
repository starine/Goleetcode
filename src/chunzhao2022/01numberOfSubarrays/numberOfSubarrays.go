package main

import "fmt"

func numberOfSubarrays(nums []int, k int) (ans int) {
	count := 0
	mp := map[int]int{}
	mp[0] = 1
	for _, num := range nums {
		count += num & 1
		if count >= k {
			ans += mp[count-k]
		}
		mp[count]++
	}
	return
}

//数学方法
func numberOfSubarrays2(nums []int, k int) (ans int) {
	odd := map[int]int{}
	count := 0
	for i, num := range nums {
		if num&1 == 1 {
			count++
			odd[count] = i
		}
	}
	odd[0] = -1
	count++
	odd[count] = len(nums)
	for i := 1; i <= count-k; i++ {
		ans += (odd[i] - odd[i-1]) * (odd[i+k] - odd[i+k-1])
	}
	return
}

func main() {
	nums := []int{1, 1, 2, 1, 1}
	k := 3
	fmt.Println(numberOfSubarrays(nums, k))
}

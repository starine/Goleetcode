package main

func twoSum(numbers []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		pre := target - numbers[i]
		if _, ok := mp[pre]; ok {
			return []int{mp[pre], i + 1}
		}
		mp[numbers[i]] = i + 1
	}
	return []int{}
}

func main() {

}

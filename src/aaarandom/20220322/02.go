package main

import (
	"fmt"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//a,_ := reader.ReadString('\n')
	//var num []int
	//for _, ch := range a {
	//	c, _ := strconv.Atoi(string(ch))
	//	num = append(num,c)
	//}
	num := []int{1, 1, 0, 1, 1}
	var ret [4]int
	var add []int
	add = append(add, num[0])
	mp := map[int][]int{}
	for i := 1; i < len(num); i++ {
		add = append(add, add[i-1]+num[i])
	}
	fmt.Println(add)
	for size := len(num) - 1; size > 0; size-- {
		mp[add[size-1]] = []int{0, size - 1}
		for start := 1; start+size-1 < len(num); start++ {
			end := start + size - 1
			count := add[end] - add[start-1]
			if _, ok := mp[count]; ok {
				pre, _ := mp[count]
				ret[0] = pre[0] + 1
				ret[1] = pre[1] + 1
				ret[2] = start + 1
				ret[3] = end + 1
				fmt.Printf("%d %d %d %d", ret[0], ret[1], ret[2], ret[3])
				return
			} else {
				mp[count] = []int{start, end}
			}
		}
	}
	fmt.Printf("%d %d %d %d", ret[0], ret[1], ret[2], ret[3])

	//l1, r1 := 0, len(num)-1
	//if l1==r1 {
	//	fmt.Printf("%d %d %d %d",1,r1,2,r1+1)
	//}
	//for num[l1] != num[r1] {
	//	l1++
	//}
	//l2, r2 := 0, len(num)-1
	//for num[l2] != num[r2] {
	//	l2--
	//}
	//if r2-l2 >= r1-l1 {
	//	fmt.Printf("%d %d %d %d",l2+1,r2,l2+2,r2+1)
	//}else {
	//	fmt.Printf("%d %d %d %d",l1+1,r1,l1+2,r1+1)
	//}
}

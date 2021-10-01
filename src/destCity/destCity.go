package main

import "fmt"

/*
1436. 旅行终点站：https://leetcode-cn.com/problems/destination-city/
*/

/*
我的解答
- 首先定义一个城市map。
- 遍历每一条path，在map中让 cityA 的 key - 1，让 cityB 的 key + 1。
- 终点只可能在cityB中出现一次，所以终点的key值为1，起点的key值为-1，其他中间城市的key为0。
- 这样做不仅可以找出终点也可以找出起点。
*/
func destCity(paths [][]string) string {
	countryMap := make(map[string]int8)
	for i := range paths {
		countryMap[paths[i][1]] += 1
		countryMap[paths[i][0]] -= 1
	}
	for k, v := range countryMap {
		if v == 1 {
			return k
		}
	}
	return ""
}

/*
官方解答
终点站不会出现在cityA中，我们可以遍历cityB，返回不在cityA中的城市，即为答案
先将cityA存于一哈希表中，然后遍历cityB的元素是否在cityA的哈希表中
*/
//func destCity(paths [][]string) string {
//	citiesA := map[string]bool{}
//	for _, path := range paths {
//		citiesA[path[0]] = true
//	}
//	for _, path := range paths {
//		if !citiesA[path[1]] {
//			return path[1]
//		}
//	}
//	return ""
//}

func main() {
	paths := [][]string{
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	}
	result := destCity(paths)
	fmt.Println(result)
}

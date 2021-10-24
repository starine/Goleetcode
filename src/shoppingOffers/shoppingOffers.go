package main

import "fmt"

/*
638. 大礼包 https://leetcode-cn.com/problems/shopping-offers/

有 n 件在售的物品
price[i] 是第 i 件物品的价格
needs[i] 是需要购买第 i 件物品的数量
special[i][j] 表示第 i 个大礼包中内含第 j 件物品的数量
special[i][n] （也就是数组中的最后一个整数）为第 i 个大礼包的价格
返回满足购物清单所需花费的最低价格
*/

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)

	// 过滤不需要计算的大礼包，只保留需要计算的大礼包
	filterSpecial := [][]int{}
	for _, s := range special {
		totalCount, totalPrice := 0, 0
		for i, c := range s[:n] {
			totalCount += c
			totalPrice += c * price[i]
		}
		if totalCount > 0 && totalPrice > s[n] {
			filterSpecial = append(filterSpecial, s)
		}
	}

	// 记忆化搜索计算满足购物清单所需花费的最低价格
	dp := map[string]int{}
	var dfs func([]byte) int
	dfs = func(curNeeds []byte) (minPrice int) {
		if res, has := dp[string(curNeeds)]; has {
			return res
		}
		for i, p := range price {
			minPrice += int(curNeeds[i]) * p // 不购买任何大礼包，原价购买购物清单中的所有物品
		}
		nextNeeds := make([]byte, n)
	outer:
		for _, s := range filterSpecial {
			for i, need := range curNeeds {
				if need < byte(s[i]) { // 不能购买超出购物清单指定数量的物品
					continue outer
				}
				nextNeeds[i] = need - byte(s[i])
			}
			minPrice = min(minPrice, dfs(nextNeeds)+s[n])
		}
		dp[string(curNeeds)] = minPrice
		return
	}

	curNeeds := make([]byte, n)
	for i, need := range needs {
		curNeeds[i] = byte(need)
	}
	return dfs(curNeeds)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	price := []int{2, 5}
	special := [][]int{{3, 0, 5}, {1, 2, 10}}
	needs := []int{3, 2}
	fmt.Println(shoppingOffers(price, special, needs))
}

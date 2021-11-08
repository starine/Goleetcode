package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
299. 猜数字游戏 https://leetcode-cn.com/problems/bulls-and-cows/
*/

func getHint(secret string, guess string) string {
	var bulls, cows int
	var cntS, cntG [10]int
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			cntS[secret[i]-'0']++
			cntG[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(cntS[i], cntG[i])
	}
	ret := strings.Builder{}
	ret.WriteString(strconv.Itoa(bulls))
	ret.WriteString("A")
	ret.WriteString(strconv.Itoa(cows))
	ret.WriteString("B")
	return ret.String()
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	secret := "1807"
	guess := "7810"
	fmt.Println(getHint(secret, guess))
}

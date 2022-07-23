/*
* @Author: starine
* @Date:   2022/7/21 11:37
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// “192.168.16.20”， “192.168.69.55”， “255.255.240.0”

func isSubset(ip1, ip2 string, mask string) bool {
	num1 := ip2Int(ip1)
	num2 := ip2Int(ip2)
	m := ip2Int(mask)
	if num1&m != num2&m {
		return false
	}
	return true
}

func ip2Int(s string) int {
	arr := strings.Split(s, ".")
	res := 0
	for i, str := range arr {
		val, _ := strconv.Atoi(str)
		res |= val << (24 - 8*i)
	}
	return res
}
func main() {
	s1 := "192.168.16.20"
	s2 := "192.168.69.55"
	mask := "255.255.240.0" //11101111
	fmt.Println(isSubset(s1, s2, mask))
	s3 := "192.168.17.20"
	fmt.Println(isSubset(s1, s3, mask))
}

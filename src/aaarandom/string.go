package main

import "fmt"

func main() {
	//字符串是不可变的,必须先将字符串转换成字节切片
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"
	fmt.Printf("s : %v, s2 : %v\n", s, s2)
	//StringIteration()
}
func StringIteration() {
	s := "中国"
	for i, v := range s {
		fmt.Printf("index : %d, value: %c\n", i, v)
	}
}

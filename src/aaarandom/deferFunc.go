package main

import "fmt"

func deferfunc1() {
	i := 0
	defer fmt.Println(i) //输出0
	i++
	defer fmt.Println(i) //输出1
	return
}
func deferfunc2() {
	i := 0
	defer func() {
		fmt.Println(i) //输出0
	}()
	i++
	defer func() {
		fmt.Println(i) //输出0
	}()
	return
}

func main() {
	deferfunc2()
}

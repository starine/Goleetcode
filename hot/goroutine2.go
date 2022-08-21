/*
* @Author: starine
* @Date:   2022/8/21 22:54
 */
//2.两个协程交替打印 1-100 数字的奇数和偶数

package main

import (
	"fmt"
	"sync"
)

var ch1, ch2 = make(chan struct{}), make(chan struct{})
var wg sync.WaitGroup

//打印奇数
func odd() {
	defer wg.Done()
	for i := 1; i < 100; i += 2 {
		<-ch1
		fmt.Println("goroutine1:",i)
		ch2 <- struct{}{}
	}
	<-ch1 //读走最后一个从协程2传入的，防止死锁
}

//打印偶数
func even() {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		<-ch2
		fmt.Println("goroutine2:",i)
		ch1 <- struct{}{}
	}
}

func main() {
	wg.Add(2)
	go odd()
	go even()
	ch1 <- struct{}{}
	wg.Wait()
}
/*
* @Author: starine
* @Date:   2022/8/21 22:50
 */
//1.交替打印数字字母
//12AB34CD56EF......2526YZ

package main

import (
	"fmt"
	"sync"
)

var numChan, strChan = make(chan struct{}), make(chan struct{})
var wg1 sync.WaitGroup

//打印数字
func printNum() {
	defer wg1.Done()
	for i := 1; i <= 26; i += 2 {
		<-numChan                  // 阻塞直到字母被打印后 numChan写入
		fmt.Printf("%v%v", i, i+1) /// 打印数字
		strChan <- struct{}{}      // strChan写入，打印字母的协程的strChan取出，才会打印字母
	}
	<-numChan
}

//打印字母
func printStr() {
	defer wg1.Done()
	for i := 65; i <= 90; i += 2 {
		<-strChan
		fmt.Printf("%v%v", string(rune(i)), string(rune(i+1)))
		numChan <- struct{}{}
	}
}

func main() {
	wg1.Add(2)
	go printNum()
	go printStr()
	numChan <- struct{}{}
	wg1.Wait()
}

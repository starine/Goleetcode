/*
* @Author: starine
* @Date:   2022/8/22 00:01
 */
//3. N个协程交替打印1-100
package main

import "fmt"
//启动N个协程，共用一个外部变量计数器，计数器范围是1到100
//开启N个有缓冲chan，chans[i]塞入数据代表协程i可以进行打印了，打印的数字就是计数器的数
//协程i一直阻塞，直到chan[i]通道有数据可以拉，才打印
func main() {
	goroutineNum := 5
	var chanSlice []chan int
	exitChan := make(chan int)

	for i := 0; i < goroutineNum; i++ { //创建N个channel
		chanSlice = append(chanSlice, make(chan int, 1))
	}

	num := 1
	j := 0
	for i := 0; i < goroutineNum; i ++ { //创建N个协程
		go func(i int) {
			for {
				<-chanSlice[i] //循环阻塞等待
				if num > 100 {
					exitChan <- 1
					break
				}

				fmt.Println(fmt.Sprintf("gorutine%v:%v", i, num))
				num++

				if j == goroutineNum-1 { //j来控制启动哪个协程来打印
					j = 0
				}else {
					j ++
				}
				chanSlice[j] <- 1
			}
		}(i)
	}
	chanSlice[0] <- 1

	select {
	case <-exitChan:
		fmt.Println("end")
	}
}

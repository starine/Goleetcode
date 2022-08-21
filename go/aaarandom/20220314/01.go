package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = str[2:]
	var sum int
	for _, ch := range str {
		ch_h, err := strconv.ParseInt(string(ch), 16, 8)
		if err != nil {
			break
		}
		for ; ch_h > 0; ch_h /= 2 {
			b := ch_h % 2
			if b == 1 {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

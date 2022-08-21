package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	strio1()
	strio2()
}

//输入有两行，第一行n
//第二行是n个字符串，字符串之间用空格隔开
//输出一行排序后的字符串，空格隔开，无结尾空格
func strio1() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	for sc.Scan() {
		strs := strings.Split(sc.Text(), " ")
		sort.Strings(strs)
		fmt.Println(strings.Join(strs, " "))
	}
}

//多个测试用例，每个测试用例一行。
//每行通过空格隔开，有n个字符，n＜100
//对于每组测试用例，输出一行排序过的字符串，每个字符串通过空格隔开
func strio2() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		strs := strings.Split(sc.Text(), " ")
		sort.Strings(strs)
		fmt.Println(strings.Join(strs, " "))
	}
}

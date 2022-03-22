package main

//推荐使用sc := bufio.NewScanner(os.Stdin)
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input()
	input1()
	input2()
}

//输入包括两个正整数a,b(1 <= a, b <= 1000),输入数据包括多组。
//输出a+b的结果
func input1() {
	var a, b int
	for {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err == io.EOF {
			break
		}
		fmt.Println(a + b)
	}
} //5ms, 812k
func input() {
	var num, res int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		for _, str := range strings.Split(sc.Text(), " ") {
			num, _ = strconv.Atoi(str)
			res += num
		}
		fmt.Println(res)
		res = 0
	}
} //3ms, 976k

//输入第一行包括一个数据组数t(1 <= t <= 100)
//接下来每行包括两个正整数a,b(1 <= a, b <= 1000)
//输出a+b的结果
func input2() {
	var t, a, b int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(a + b)
	}
} //5ms,1200kb
func input22() {
	var t, num, sum int
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	t, _ = strconv.Atoi(sc.Text())
	for i := 0; i < t; i++ {
		sc.Scan()
		for _, str := range strings.Fields(sc.Text()) {
			num, _ = strconv.Atoi(str)
			sum += num
		}
		fmt.Println(sum)
		sum = 0
	}
} //3ms,1068kb

func input3() {
	var a, b int
	for {
		fmt.Scanf("%d %d", &a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
} //6ms, 1064kb
func input33() {
	var a, b int
	var nums []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		nums = strings.Split(sc.Text(), " ")
		a, _ = strconv.Atoi(nums[0])
		b, _ = strconv.Atoi(nums[1])
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
} //3ms,972kb

//输入数据包括多组。
//每组数据一行,每行的第一个整数为整数的个数n(1 <= n <= 100), n为0的时候结束输入。
//接下来n个正整数,即需要求和的每个正整数。
func input4() {
	var n, tmp int
	for {
		ans := 0
		fmt.Scanf("%d", &n)
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &tmp)
			ans += tmp
		}
		fmt.Println(ans)
	}
} //137ms,2492kb
func input44() {
	var sum int
	var nums []int
	var str []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		str = strings.Fields(sc.Text())
		if len(str) == 1 {
			break
		}
		for _, num := range str {
			n, _ := strconv.Atoi(num)
			sum += n
			nums = append(nums, n)
		}
		sum -= nums[0]
		fmt.Println(sum)
		sum = 0
		nums = nums[0:0]
		str = str[0:0]
	}
} //5ms,1776kb

//输入数据有多组, 每行表示一组输入数据。
//每行不定有n个整数，空格隔开。(1 <= n <= 100)。
//输出每组的和
func input7() {
	var num, sum int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		for _, str := range strings.Fields(sc.Text()) {
			num, _ = strconv.Atoi(str)
			sum += num
		}
		fmt.Println(sum)
		sum = 0
	}
}

package main

import "fmt"

/*
476. 数字的补数 https://leetcode-cn.com/problems/number-complement/
*/
/*
位运算
1.首先找到 num 二进制表示最高位的那个 1，再将这个 1 以及更低的位进行取反。

	如果 num 二进制表示最高位的 1 是第 i(0≤i≤30) 位，那么一定有：
		2^i < num < 2^{i+1}
	因此我们可以使用一次遍历，在 [0, 30][0,30] 中找出 ii 的值。

2.遍历 num 的第 0∼i 个二进制位，将它们依次进行取反。

	可以用更高效的方式，构造掩码 mask = 2^{i+1} - 1
	它是一个 i+1 位的二进制数，并且每一位都是 1。我们将 num 与 mask 进行异或运算，0^1=1, 1^1=0, 即可得到答案。
*/
func findComplement(num int) int {
	highBit := 0
	for i := 1; i <= 30; i++ {
		if num < 1<<i {
			break
		}
		highBit = i
	}
	mask := 1<<(highBit+1) - 1 //构造掩码
	//fmt.Printf("%b\n",mask)
	return num ^ mask
}

func main() {
	num := 5
	//fmt.Printf("%b\n",num)
	fmt.Println(findComplement(num))
}

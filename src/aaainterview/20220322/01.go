package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nk := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])
	var arr [][]int
	for i := 0; i < n; i++ {
		sc.Scan()
		var row []int
		for _, s := range strings.Split(sc.Text(), " ") {
			c, _ := strconv.Atoi(s)
			row = append(row, c)
		}
		arr = append(arr, row)
	}
	var res [][]int
	for i := 0; i < n; i++ {
		var row1 []int
		for _, num := range arr[i] {
			for j := 0; j < k; j++ {
				row1 = append(row1, num)
			}
		}
		for m := 0; m < k; m++ {
			res = append(res, row1)
		}
	}
	for i := 0; i < n*k; i++ {
		for _, num := range res[i] {
			fmt.Printf("%d ", num)
		}
		fmt.Print("\n")
	}
}

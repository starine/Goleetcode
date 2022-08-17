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
	nm := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])
	var arr [][]int
	for i := 0; i < n; i++ {
		sc.Scan()
		var row []int
		for _, num := range strings.Split(sc.Text(), " ") {
			r, _ := strconv.Atoi(num)
			row = append(row, r)
		}
		arr = append(arr, row)
	}
	fmt.Println(arr)

	l2r := make([][]int, n)
	for i := 0; i < n; i++ {
		l2r[i] = make([]int, m)
	}
	u2b := make([][]int, n)
	for i := 0; i < n; i++ {
		u2b[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		l2r[i][0] = arr[i][0]
	}
	for j := 0; j < m; j++ {
		u2b[0][j] = arr[0][j]
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			l2r[i][j] = l2r[i][j-1] + arr[i][j]
		}
	}
	for j := 0; j < m; j++ {
		for i := 1; i < n; i++ {
			u2b[i][j] = u2b[i-1][j] + arr[i][j]
		}
	}

	var score int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				if l2r[i][j] != 0 {
					score++
				}
				if u2b[i][j] != 0 {
					score++
				}
				if i != n-1 && u2b[n-1][j] > u2b[i][j] {
					score++
				}
				if j != m-1 && l2r[i][m-1] > l2r[i][j] {
					score++
				}
			}
		}
	}
	fmt.Println(score)
}

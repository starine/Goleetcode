/*
* @Author: starine
* @Date:   2022/7/24 10:45
 */

package main

func equalPairs(grid [][]int) int {
	n := len(grid)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if grid[i][k] != grid[k][j] {
					break
				}
				res++
			}
		}
	}
	return res
}
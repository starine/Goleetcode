package main

func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	if row == 0 { //空矩阵的情况
		return 0
	}
	res := 0 //记录岛屿数
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' { //遍历到1的情况
				res++
				DFS(grid, i, j)
			}
		}
	}
	return res
}
func DFS(grid [][]byte, r, c int) { //深度优先遍历
	row, col := len(grid), len(grid[0])
	grid[r][c] = '0' //置为0
	//相邻四个⽅向深度遍历
	if r-1 >= 0 && grid[r-1][c] == '1' {
		DFS(grid, r-1, c)
	}
	if r+1 < row && grid[r+1][c] == '1' {
		DFS(grid, r+1, c)
	}
	if c-1 >= 0 && grid[r][c-1] == '1' {
		DFS(grid, r, c-1)
	}
	if c+1 < col && grid[r][c+1] == '1' {
		DFS(grid, r, c+1)
	}
}

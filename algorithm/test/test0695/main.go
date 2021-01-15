package main

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	maxNum := 0
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				maxNum = max(maxNum, dfs(grid, rows, cols, i, j))
			}
		}
	}
	return maxNum
}

func dfs(grid [][]int, rows, cols, i, j int) int {
	if i < 0 || j < 0 || i >= rows || j >= cols {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 0
	up := dfs(grid, rows, cols, i-1, j)
	down := dfs(grid, rows, cols, i+1, j)
	left := dfs(grid, rows, cols, i, j-1)
	right := dfs(grid, rows, cols, i, j+1)
	return up + down + left + right + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

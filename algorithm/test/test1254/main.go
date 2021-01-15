package main

func main() {

}

func closedIsland(grid [][]int) int {
	res := 0
	row := len(grid)
	col := len(grid[0])
	if row == 0 && col == 0 {
		return res
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if grid[i][j] == 0 {
				if dfs(row, col, i, j, grid) {
					res++
				}
			}
		}
	}
	return res
}

func dfs(row, col, i, j int, grid [][]int) bool {
	if i < 0 || j < 0 || i >= row || j >= col {
		return false
	}
	if grid[i][j] == 1 {
		return true
	}
	grid[i][j] = 1
	up := dfs(row, col, i-1, j, grid)
	down := dfs(row, col, i+1, j, grid)
	left := dfs(row, col, i, j-1, grid)
	right := dfs(row, col, i, j+1, grid)
	if up && down && left && right {
		return true
	}
	return false
}

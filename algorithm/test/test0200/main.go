package main

import "fmt"

func main() {
	grid := [][]byte{[]byte{1, 1, 0, 0, 0}, []byte{1, 1, 0, 0, 0}, []byte{0, 0, 1, 0, 0}, []byte{0, 0, 0, 1, 1}}
	for i := 0; i < len(grid); i++ {
		fmt.Println()
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}
	}
	res := numIslands(grid)
	fmt.Println(res)
	for i := 0; i < len(grid); i++ {
		fmt.Println()
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}
	}
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if grid[i][j] != 1 {
		return
	}
	grid[i][j] = 0
	if i > 0 {
		dfs(grid, i-1, j)
	}
	if j > 0 {
		dfs(grid, i, j-1)
	}
	if i < len(grid)-1 {
		dfs(grid, i+1, j)
	}
	if j < len(grid[i])-1 {
		dfs(grid, i, j+1)
	}
}

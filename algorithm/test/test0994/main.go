package main

import "fmt"

func main() {
	grid := [][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	if len(grid[0]) == 0 {
		return -1
	}
	// 找出所有坏掉的orange
	bad := make([][]int, 0, len(grid)*len(grid[0]))
	good := make([][]int, 0, len(grid)*len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				bad = append(bad, []int{i, j})
			}
			if grid[i][j] == 1 {
				good = append(good, []int{i, j})
			}
		}
	}
	if len(good) == 0 {
		return 0
	}
	count := 0
	if len(bad) == 0 {

		return 0
	}
	length := len(bad)
	tmp := length
	for i := 0; i < length; i++ {
		b := bad[i]
		if b[0] > 0 && grid[b[0]-1][b[1]] == 1 {
			length++
			grid[b[0]-1][b[1]] = 2
			bad = append(bad, []int{b[0] - 1, b[1]})
		}
		if b[0] < len(grid)-1 && grid[b[0]+1][b[1]] == 1 {
			length++

			grid[b[0]+1][b[1]] = 2
			bad = append(bad, []int{b[0] + 1, b[1]})
		}
		if b[1] > 0 && grid[b[0]][b[1]-1] == 1 {
			length++

			grid[b[0]][b[1]-1] = 2
			bad = append(bad, []int{b[0], b[1] - 1})
		}
		if b[1] < len(grid[0])-1 && grid[b[0]][b[1]+1] == 1 {
			length++
			grid[b[0]][b[1]+1] = 2
			bad = append(bad, []int{b[0], b[1] + 1})
		}
		if tmp == i+1 {
			tmp = length
			count++
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return count - 1
}

func printGrid(grid [][]int) {
	for _, n := range grid {
		fmt.Println(n)
	}
	fmt.Println()
}

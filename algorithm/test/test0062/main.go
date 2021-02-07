package main

func main() {

}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	cells := make([][]int, m)
	for index, _ := range cells {
		cells[index] = make([]int, n)
	}
	for i, rows := range cells {
		for j, _ := range rows {
			if i == 0 && j == 0 {
				cells[i][j] = 1
				continue
			}
			cells[i][j] = cells[i-1][j] + cells[i][j-1]
		}
	}
	return cells[m-1][n-1]
}

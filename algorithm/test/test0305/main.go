package main

import "fmt"

func main() {
	positions := [][]int{[]int{0, 1}, []int{0, 0}}
	m, n := 1, 2
	res := numIslands2(m, n, positions)
	fmt.Println(res)
}

func numIslands2(m int, n int, positions [][]int) []int {
	count := 0
	var record []int
	island := make([][]bool, m)
	for index, _ := range island {
		island[index] = make([]bool, n)
	}
	for _, v := range positions {
		island[v[0]][v[1]] = true
		exist := false
		if v[0] > 0 && island[v[0]-1][v[1]] {
			exist = true
		}
		if v[1] > 0 && island[v[0]][v[1]-1] {
			exist = true
		}

		if v[0] < m-1 && island[v[0]+1][v[1]] {
			exist = true
		}

		if v[1] < n-1 && island[v[0]][v[1]+1] {
			exist = true
		}
		if !exist {
			count++
		}
		record = append(record, count)
	}
	return record
}

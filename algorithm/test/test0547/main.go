package main

import "fmt"

func main() {
	args := [][]int{[]int{1, 1, 0}, []int{1, 1, 0}, []int{0, 0, 1}}
	res := findCircleNumV3(args)
	fmt.Println(res)
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]int, len(isConnected))
	count := 0
	for i := 0; i < len(visited); i++ {
		if visited[i] == 0 {
			count++
			dfs(isConnected, visited, i)
		}
	}
	return count
}

func dfs(isConnected [][]int, visited []int, i int) {
	for j := 0; j < len(visited); j++ {
		if visited[j] == 0 && isConnected[i][j] == 1 {
			visited[j] = 1
			dfs(isConnected, visited, j)
		}
	}
}

func findCircleNumV2(isConnected [][]int) int {
	visited := make([]int, len(isConnected))
	count := 0
	record := make([]int, 0)
	for i := 0; i < len(visited); i++ {
		if visited[i] == 1 {
			continue
		}
		count++
		record = append(record, i)
		for len(record) != 0 {
			root := record[0]
			visited[record[0]] = 1
			for j := 0; j < len(visited); j++ {
				if visited[j] == 0 && isConnected[root][j] == 1 {
					record = append(record, j)
				}
			}
			record = record[1:]
		}
	}
	return count
}

func findCircleNumV3(isConnected [][]int) int {
	parent := make([]int, len(isConnected))
	for index, _ := range parent {
		parent[index] = index
	}
	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected); j++ {
			union(parent, i, j)
		}
	}
	count := 0
	for index, v := range parent {
		if v == index {
			count++
		}
	}
	return count
}

func union(parent []int, x, y int) {
	xp := find(parent, x)
	yp := find(parent, y)
	parent[yp] = xp
}

func find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return find(parent, parent[i])
}

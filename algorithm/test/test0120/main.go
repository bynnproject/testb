package main

import (
	"fmt"
	"math"
)

func main() {
	res := minimumTotal([][]int{[]int{1}, []int{2, 3}})
	fmt.Println(res)
}

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for index, target := range triangle {
		fmt.Println(index, len(target))
		dp[index] = make([]int, len(target))
		if index == 0 {
			dp[index][0] = triangle[0][0]
			continue
		}
		tmpLength := len(target)
		for i, _ := range target {
			if i == 0 {
				dp[index][i] = dp[index-1][i] + target[i]
				continue
			}
			if i == tmpLength-1 {
				dp[index][i] = dp[index-1][i-1] + target[i]
				continue
			}
			dp[index][i] = min(dp[index-1][i-1], dp[index-1][i]) + target[i]
		}
	}
	minNum := math.MaxInt32
	for _, num := range dp[length-1] {
		minNum = min(minNum, num)
	}
	return minNum
}

//func minimumTotal(triangle [][]int) int {
//	if len(triangle) == 1 {
//		return triangle[0][0]
//	}
//	return dfs(triangle, 0, 0, 0)
//}

func dfs(triangle [][]int, deep, position, record int) int {
	fmt.Println(deep, position, record, record+triangle[deep][position])
	if len(triangle)-1 == deep {
		return record + triangle[deep][position]
	}

	left := dfs(triangle, deep+1, position, record+triangle[deep][position])
	if position == len(triangle[deep]) {
		return left
	}
	right := dfs(triangle, deep+1, position+1, record+triangle[deep][position])
	return min(left, right)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

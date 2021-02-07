package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findLength(
		[]int{0, 1, 1, 1, 1},
		[]int{1, 0, 1, 0, 1}))
}

func findLength(A []int, B []int) int {
	ALength := len(A)
	BLength := len(B)
	if ALength == 0 || BLength == 0 {
		return 0
	}
	dp := make([][]int, ALength+1)
	for index, _ := range dp {
		dp[index] = make([]int, BLength+1)
	}
	max := 0
	for i := ALength - 1; i >= 0; i-- {
		for j := BLength - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	math.MinInt32
	return max
}

func find(A, B []int) int {
	for windowLength := len(A); windowLength > 0; windowLength-- {
		for right := len(A) - 1; right > 0; right-- {
			if right-windowLength < 0 {
				break
			}

		}
	}
	return 0

}

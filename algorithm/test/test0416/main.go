package main

import "fmt"

func main() {
	fmt.Println(canPartitionV2([]int{1, 5, 11, 5}))
}

func canPartitionV2(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}
	var (
		sum int
		max int
	)
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if 0 != sum%2 {
		return false
	}
	sum /= 2
	if max > sum {
		return false
	}
	dp := make([][]bool, length)

	for i, _ := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < length; i++ {
		for j := 1; j < sum+1; j++ {
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
				continue
			}
			dp[i][j] = dp[i-1][j]
		}
	}
	return dp[length-1][sum]
}

func printDp(dp [][]bool) {
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	fmt.Println()
}

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if 0 != sum%2 {
		return false
	}
	sum /= 2
	for length := 1; length <= len(nums)/2+1; length++ {
		for j := 0; j < len(nums)-length; j++ {
			if record(nums, j, length, sum) {
				fmt.Println(j, length, sum)
				return true
			}
		}
	}
	return false
}

func record(nums []int, position, length, val int) bool {
	val -= nums[position]
	length--
	if length == 0 {
		return val == 0
	}
	for i := 1; i <= len(nums)-position-length; i++ {
		if record(nums, position+i, length, val) {

			return true
		}
	}
	return false
}

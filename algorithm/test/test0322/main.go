package main

import (
	"fmt"
	"math"
)

func main()  {
	testCase := []struct {
		coins []int
		amount int
		res int
	}{
		{
			coins: []int{2},
			amount: 3,
			res : -1,
		},


	}
	for idx ,c := range testCase {
		if c.res != coinChangeV2(c.coins ,c.amount ){
			fmt.Printf("false : %d" , idx)
			return
		}else {
			fmt.Printf("ok : %d" , idx)
			return
		}
	}
}


func coinChange(coins []int, amount int) int {
	return count(coins , 0, amount)
}

var maxNum = math.MaxInt32

func count(coins []int, coinsIdx , amount int) int {
	if amount == 0 {
		return 0
	}

	if coinsIdx >= len(coins) || amount < 0 {
		return -1
	}
	num := amount / coins[coinsIdx]
	minCoins := maxNum
	for i := 0 ; i <= num ; i++ {
		res := count(coins , coinsIdx + 1 , amount - i * coins[coinsIdx])
		if res != -1 {
			minCoins = min(minCoins , res + i)
		}
	}
	if minCoins == maxNum {
		return -1
	}
	return minCoins
}

func min(i , j int) int  {
	if i < j {
		return i
	}
	return j
}


func coinChangeV2(coins []int, amount int) int {
	dp := make([]int , amount + 1)
	dp[0] = 0
	for i := 1 ; i <= amount ; i++ {
		dp[i] = i + 1
		for _ , c := range coins {
			if i < c {
				continue
			}
			if dp[i - c] == i - c + 1 {
				continue
			}
			dp[i] = min(dp[i] , dp[i - c] + 1)
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}


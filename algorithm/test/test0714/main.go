package main

func main() {

}

func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}
	for i, price := range prices {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -price - fee
			continue
		}
		dp[i][0] = max(dp[i-1][1]+price, dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-price-fee, dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func maxProfitV2(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	sell, buy := 0, prices[0]-fee
	for i := 1; i < len(prices); i++ {
		sell, buy = max(sell, buy+prices[i]), max(sell-prices[i]-fee, buy)
	}
	return sell
}

func maxProfitV3(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	sell, buy := 0, prices[0]-fee
	for i := 1; i < len(prices); i++ {
		sell, buy = max(sell, buy+prices[i]), max(sell-prices[i]-fee, buy)
	}
	return sell
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

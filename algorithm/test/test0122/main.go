package main

func main()  {
	
}

func maxProfit(prices []int) int  {
	n := len(prices)
	dp := make([][2]int , n)
	dp[0][1] = -prices[0]
	for i := 1 ; i < n ; i++ {
		dp[i][0] = max(dp[i -1][0] , dp[i - 1][1] + prices[i])
		dp[i][1] = max(dp[i - 1][1] , dp[i -1][0] - prices[i])
	}
	return dp[n - 1][0]
}


func maxProfitV2(prices []int) int  {
	res := 0
	for i := 1 ; i < len(prices); i++ {
		res += max( 0 , prices[i] - prices[i -1])
	}
	return res
}





func max( i , j int) int {
	if i > j {
		return i
	}
	return j
}
package main

import "fmt"

func main() {
	res := wordBreak("a", []string{"a"})
	fmt.Println(res)
}

func wordBreak(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)
	for _, w := range wordDict {
		wordDictMap[w] = true
	}
	length := len(s)
	dp := make([]bool, length+1)
	dp[0] = true
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && wordDictMap[s[j:i+1]] {
				dp[i+1] = true
				break
			}
		}
	}
	return dp[length]
}

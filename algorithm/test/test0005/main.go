package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {
	length := len(s)
	start, end := 0, 0
	for i := 0; i < length; i++ {
		left, right := aroundAsCenter(s, i, i)
		fmt.Println(start, end, left, right)
		if right-left > end-start {
			start, end = left, right
		}
		left, right = aroundAsCenter(s, i, i+1)
		fmt.Println(start, end, left, right)
		if right-left > end-start {
			start, end = left, right
		}
	}

	return s[start : end+1]
}

func aroundAsCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

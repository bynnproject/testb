package main

import "fmt"

func main() {
	testCase := []struct {
		param string
		res   int
	}{
		{
			param: "abcabcbb",
			res:   3,
		},
		{
			param: "bbbbb",
			res:   1,
		},
		{
			param: "pwwkew",
			res:   3,
		},
		{
			param: "a",
			res:   0,
		},
	}
	for index, c := range testCase {
		res := lengthOfLongestSubstring(c.param)
		fmt.Printf("case index %d : param %s res %d \n", index, c.param, c.res)
		fmt.Printf("result %d \n", res)
		fmt.Printf("%t \n", res == c.res)
	}

}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	var (
		left   int
		right  int
		maxLen int
	)
	bs := []uint8(s)
	recordMap := make(map[uint8]int)
	recordMap[bs[right]] = 0
	maxLen = 1
	for right < length-1 {
		right++
		if v, ok := recordMap[bs[right]]; ok {
			for left != v+1 {
				delete(recordMap, bs[left])
				left++
			}
		}
		recordMap[bs[right]] = right
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

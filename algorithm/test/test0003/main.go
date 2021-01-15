package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("aaabcefddcaq"))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxNum := 0
	tmp := 0
	length := len(s)
	left := 0
	for right := 0; right < length; right++ {
		index, ok := m[s[right]]
		tmp++
		fmt.Printf("index: %d\n", index)
		if !ok {
			fmt.Printf("add %d \n", right)
			m[s[right]] = right
			continue
		}
		maxNum = max(maxNum, tmp-1)
		for left != index+1 {
			fmt.Printf("delete %d\n", left)
			delete(m, s[left])
			left++
			tmp--
		}
		m[s[right]] = right
		fmt.Printf("tmp %d\n", tmp)
	}
	maxNum = max(maxNum, tmp)
	return maxNum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

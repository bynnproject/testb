package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
}

func largestNumber(nums []int) string {

	numStrs := make([]string, 0, len(nums))
	for _, num := range nums {
		numStrs = append(numStrs, strconv.FormatInt(int64(num), 10))
	}
	sort.Slice(numStrs, func(i, j int) bool {
		a := numStrs[i] + numStrs[j]
		b := numStrs[j] + numStrs[i]
		if a > b {
			return true
		}
		return false
	})
	s := strings.Join(numStrs, "")
	return s
}

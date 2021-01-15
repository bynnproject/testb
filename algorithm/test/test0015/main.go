package main

import (
	"fmt"
	"sort"
)

func main() {
	args := []int{1, 2, -2, -1}
	final := threeSum(args)
	fmt.Println(final)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	length := len(nums)
	for first := 0; first < length; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := 0 - nums[first]
		third := length - 1
		for second := first + 1; second < length-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] != target {
				third--
			}
			if second >= third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ans

}

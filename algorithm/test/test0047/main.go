package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	fmt.Println(nums)
	helper(nums, 0, len(nums)-1, &res)
	return res
}

func helper(nums []int, begin, end int, res *[][]int) {
	if begin == end {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	//if begin > 0 && nums[begin-1] == nums[begin] {
	//	return
	//}
	for i := begin; i <= end; i++ {
		if i > 0 && nums[i] == nums[i-1] && begin < i {
			continue
		}
		nums[begin], nums[i] = nums[i], nums[begin]
		helper(nums, begin+1, end, res)
		nums[i], nums[begin] = nums[begin], nums[i]
	}
}

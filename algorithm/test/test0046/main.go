package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	var res [][]int
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
	for i := begin; i <= end; i++ {
		nums[begin], nums[i] = nums[i], nums[begin]
		helper(nums, begin+1, end, res)
		nums[i], nums[begin] = nums[begin], nums[i]
	}
}

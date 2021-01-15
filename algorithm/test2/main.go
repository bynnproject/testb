package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("final", rob2(nums))
}

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func rob2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	fmt.Println(rob(nums[:length-1]))
	fmt.Println(rob(nums[1:]))
	return max(rob(nums[:length-1]), rob(nums[1:]))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

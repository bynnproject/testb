package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

func findDuplicate(nums []int) int {
	for index, _ := range nums {
		for nums[index] != index {
			if nums[nums[index]] == nums[index] {
				return nums[index]
			}
			n := nums[index]
			nums[index], nums[n] = nums[n], nums[index]
		}
	}
	return -1
}

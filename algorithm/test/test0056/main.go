package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{[]int{1, 4}, []int{2, 3}}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			if intervals[i][1] > intervals[j][1] {
				return false
			}
			return true
		}
		if intervals[i][0] > intervals[j][0] {
			return false
		}
		return true
	})
	var newIntervals [][]int
	length := len(intervals)
	for i := 0; i < length; i++ {
		if i == length-1 {
			newIntervals = append(newIntervals, intervals[i])
			break
		}
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][0] = intervals[i][0]
			if intervals[i][1] > intervals[i+1][1] {
				intervals[i+1][1] = intervals[i][1]
			}
			continue
		}
		newIntervals = append(newIntervals, intervals[i])
	}

	return newIntervals
}

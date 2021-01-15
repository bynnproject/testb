package main

import "fmt"

func main() {
	args := []int{3, 4, 6, 1, 2, 5, 9, 8, 10, 7}
	mergeSort(args)
	fmt.Println(args)
}

func bubbleSort(args []int) {
	exchange := false
	for i := 0; i < len(args); i++ {
		exchange = false
		for j := 0; j < len(args)-i-1; j++ {
			if args[j] > args[j+1] {
				exchange = true
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
		if !exchange {
			return
		}
	}
}

func selectSort(args []int) {
	for i := 0; i < len(args)-1; i++ {
		pos := i
		for j := i + 1; j < len(args); j++ {
			if args[pos] > args[j] {
				pos = j
			}
		}
		args[i], args[pos] = args[pos], args[i]
	}
}

func insertSort(args []int) {
	for i := 1; i < len(args); i++ {
		for j := i; j > 0 && args[j] < args[j-1]; j-- {
			args[j], args[j-1] = args[j-1], args[j]
		}
	}
}

func shellSort(args []int) {
	for gap := len(args) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(args); i += gap {
			for j := i; j > 0 && args[j] > args[j-gap]; j = j - gap {
				args[j], args[j-gap] = args[j-gap], args[j]
			}
		}
	}
}

func quickSort(args []int) {
	if len(args) == 0 || len(args) == 1 {
		return
	}
	mid := args[0]
	left, right := 1, len(args)-1
	for left < right {
		if args[left] <= mid {
			left++
			continue
		}
		if args[right] >= mid {
			right--
			continue
		}
		args[left], args[right] = args[right], args[left]
	}
	args[left-1], args[0] = args[0], args[left-1]
	quickSort(args[:left])
	quickSort(args[left:])
}

func headSort(args []int) {
	for i := len(args)/2 - 1; i >= 0; i-- {
		adjustNode(args, i, len(args)-1)
	}
	for i := len(args) - 1; i > 0; i-- {
		if i == 1 && args[0] < args[i] {
			break
		}
		args[0], args[i] = args[i], args[0]
		adjustNode(args, 0, i-1)
	}
}

func adjustNode(args []int, start, end int) {
	leftNodeNum := start*2 + 1
	if leftNodeNum > end {
		return
	}
	rightNodeNum := leftNodeNum + 1
	largeChildNodeNum := leftNodeNum
	if rightNodeNum <= end && args[rightNodeNum] > args[leftNodeNum] {
		largeChildNodeNum = rightNodeNum
	}
	if args[start] >= args[largeChildNodeNum] {
		return
	}
	args[start], args[largeChildNodeNum] = args[largeChildNodeNum], args[start]
	adjustNode(args, largeChildNodeNum, end)
}

func mergeSort(args []int) {
	if len(args) == 1 {
		return
	}
	mid := len(args) / 2
	left, right := 0, mid
	mergeSort(args[:mid])
	mergeSort(args[mid:])
	tmp := make([]int, 0, len(args))
	for left < mid && right < len(args) {
		if args[left] < args[right] {
			tmp = append(tmp, args[left])
			left++
			continue
		}
		tmp = append(tmp, args[right])
		right++
	}
	if left < mid {
		tmp = append(tmp, args[left:mid]...)
	}
	if right < len(args) {
		tmp = append(tmp, args[right:]...)
	}
	for index, value := range tmp {
		args[index] = value
	}
}

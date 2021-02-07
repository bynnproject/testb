package main

import "fmt"

func main()  {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}

func dailyTemperatures(T []int) []int {
	length := len(T)
	res := make([]int , length)
	p := -1
	stack := make([]int , length)
	for i := 0; i < length; i++ {
		if p == -1 {
			p++
			stack[p] = i
			continue
		}
		if T[stack[p]] >= T[i] {
			p++
			stack[p] = i
			continue
		}
		for p != -1 && T[stack[p]] < T[i]{
			res[stack[p]] = i - stack[p]
			p--
		}
		p++
		stack[p] = i
	}
	return res
}

package main

import "fmt"

func main() {
	num := "1432219"
	k := 3
	rs := removeKdigits(num, k)
	fmt.Println(rs)
}

func removeKdigits(num string, k int) string {
	ns := []uint8(num)
	lastZeroIndex := 0
	for i := 0; i < k; i++ {
		if ns[i] == 0 {
			lastZeroIndex = i
		}
	}
	if lastZeroIndex != 0 {
		ns = ns[lastZeroIndex:]
		k = k - lastZeroIndex - 1
	}
	low := 0
	low
	for _, n := range ns {
		l
	}
	tmp := make([]uint8, 0)
	for _, n := range ns {
		if n != 0 {
			tmp = append(tmp, n)
		}
	}
	fmt.Println(tmp)
	fmt.Println(string(tmp))
	return string(ns)
}

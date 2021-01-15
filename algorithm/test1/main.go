package main

import "fmt"

//在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
//如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
//
func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	i := canCompleteCircuit(gas, cost)
	fmt.Println(i)
}

//func canCompleteCircuit(gas []int, cost []int) int {
//	if len(gas) != len(cost) {
//		return -1
//	}
//	length := len(gas)
//	for i := 0; i < length; {
//		count, sum := 0, 0
//		for j := i; count < length; {
//			sum += gas[j] - cost[j]
//			j = (j + 1) % length
//			count++
//			if sum < 0 {
//				break
//			}
//		}
//		if sum < 0 {
//			i = i + count
//		} else {
//			return i
//		}
//	}
//	return -1
//}
func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) != len(cost) {
		return -1
	}
	length := len(gas)
	for i := 0; i < length; {
		sum, count := 0, 0
		for j := i; count < length; j = (j + 1) % length {
			count++
			sum += gas[j] - cost[j]
			if sum < 0 {
				break
			}
		}

		if sum < 0 {
			i = i + count
		} else {
			return i
		}
	}
	return -1
}

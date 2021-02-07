package main

import "fmt"

func main() {
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l1.Right = l2
	l2.Right = l3
	l3.Right = l4
	res := pathSum(l1, 3)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {

	return calPathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func calPathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum -= root.Val
	if sum == 0 {
		return calPathSum(root.Left, sum) + calPathSum(root.Right, sum) + 1
	}
	return calPathSum(root.Left, sum) + calPathSum(root.Right, sum)
}

func pathSumV2(root *TreeNode, sum, trueSum int) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		if root.Val == sum {
			return 1
		}
		return 0
	}
	var count int
	if root.Val == sum {
		count++
	}
	fmt.Println(root.Val, count)
	count += pathSumV2(root.Left, sum-root.Val, trueSum)
	fmt.Println(root.Val, count)

	count += pathSumV2(root.Left, trueSum, trueSum)
	fmt.Println(root.Val, count)

	count += pathSumV2(root.Right, sum-root.Val, trueSum)
	fmt.Println(root.Val, count)

	count += pathSumV2(root.Right, trueSum, trueSum)
	fmt.Println(root.Val, count)

	fmt.Println()

	return count
}

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	sums := sumMaxValue(root)
	return max(sums[0], sums[1])
}

func sumMaxValue(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l, r := sumMaxValue(root.Left), sumMaxValue(root.Right)
	withRoot := root.Val + l[1] + r[1]
	notWithRoot := max(l[0], l[1]) + max(r[0], r[1])
	return []int{withRoot, notWithRoot}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

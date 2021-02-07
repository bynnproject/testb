package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			pre := curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Left, curr.Right = nil, curr.Left
		}
		curr = curr.Right
	}
}

package main

func main()  {
	
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var (
		record *[]int
		final *[][]int
	)

	return nil
}

func count(root *TreeNode , sum int , record *[]int , final *[][]int){
	if root.Val == sum {
		if root.Left != nil || root.Right != nil {
		}
		*record = append(*record, root.Val)
		*final = append(*final, *record)
	}
	if sum < root.Val {
		return
	}
	*record = append(*record, root.Val)
	if root.Left != nil && count(root.Left , sum - root.Val , record,final) {
		return
	}
	if root.Right != nil && count(root.Right , sum - root.Val , record,final) {
		return
	}
	return
}
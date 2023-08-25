package easy

/**
 * Definition for a binary tree node.
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	// var process func(t int, x, y bool) bool
	process := func(t int, x, y bool) bool {
		if t == 2 {
			return x || y
		} else {
			return x && y
		}
	}
	var eva func(root *TreeNode) bool
	eva = func(root *TreeNode) bool {
		if root.Left == nil {
			return root.Val == 1
		}
		return process(root.Val, eva(root.Left), eva(root.Right))
	}
	return eva(root)
}

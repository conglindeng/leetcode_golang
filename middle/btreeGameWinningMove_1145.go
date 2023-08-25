package middle

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var node *TreeNode
	var getCount func(*TreeNode) int
	getCount = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		if tn.Val == x {
			node = tn
		}
		return 1 + getCount(tn.Left) + getCount(tn.Right)
	}
	getCount(root)
	l := getCount(node.Left)
	if l > n/2 {
		return true
	}
	r := getCount(node.Right)
	if r > n/2 {
		return true
	}
	return n-l-r-1 > n/2
}

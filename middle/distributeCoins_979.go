package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		leftC, leftV := dfs(root.Left)
		rightC, rightV := dfs(root.Right)
		c := leftC + rightC + 1
		v := leftV + rightV + root.Val
		res += abs(c - v)
		return c, v
	}
	dfs(root)
	return res
}

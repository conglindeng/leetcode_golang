package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) (int, *TreeNode)
	dfs = func(root *TreeNode) (int, *TreeNode) {
		if root == nil {
			return 0, nil
		}
		i, left := dfs(root.Left)
		j, right := dfs(root.Right)
		if i < j {
			return j + 1, right
		} else if i > j {
			return i + 1, left
		}
		return i + 1, root
	}
	_, tn := dfs(root)
	return tn
}

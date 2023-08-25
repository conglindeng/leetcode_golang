package middle

import "github.com/conglindeng/leetcode/mystruct"

func pruneTree(root *mystruct.TreeNode) *mystruct.TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

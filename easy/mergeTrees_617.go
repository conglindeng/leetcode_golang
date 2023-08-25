package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	} else {
		res := &TreeNode{Val: root1.Val + root2.Val}
		res.Left = mergeTrees(root1.Left, root2.Left)
		res.Right = mergeTrees(root1.Right, root2.Right)
		return res
	}
}

func mergeTrees_bfs(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	queue1 := make([]*TreeNode, 0)
	queue2 := make([]*TreeNode, 0)
	queue1 = append(queue1, root1)
	queue2 = append(queue2, root2)

	for len(queue1) > 0 && len(queue2) > 0 {

	}

	return nil
}

package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func lowestCommonAncestor_stack(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			parent[root.Right.Val] = root
			dfs(root.Right)
		}
	}
	dfs(root)
	visit := map[int]bool{}
	for p != nil {
		visit[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visit[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

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

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue1 := make([]*TreeNode, 0)
	queue2 := make([]*TreeNode, 0)
	work := []*TreeNode{}

	merged := &TreeNode{Val: root1.Val + root2.Val}
	work = append(work, merged)
	queue1 = append(queue1, root1)
	queue2 = append(queue2, root2)
	for len(queue1) > 0 && len(queue2) > 0 {
		cur := work[0]
		work = work[1:]
		temp1 := queue1[0]
		temp2 := queue2[0]
		queue1 = queue1[1:]
		queue2 = queue2[1:]
		left1, right1 := temp1.Left, temp1.Right
		left2, right2 := temp2.Left, temp2.Right
		if left1 != nil && left2 != nil {
			newLeft := &TreeNode{Val: left1.Val + left2.Val}
			cur.Left = newLeft
			work = append(work, newLeft)
			queue1 = append(queue1, temp1.Left)
			queue2 = append(queue2, temp2.Left)
		} else if left1 != nil {
			cur.Left = left1
		} else if left2 != nil {
			cur.Left = left2
		}

		if right1 != nil && right2 != nil {
			newRight := &TreeNode{Val: right1.Val + right2.Val}
			cur.Right = newRight
			work = append(work, newRight)
			queue1 = append(queue1, temp1.Right)
			queue2 = append(queue2, temp2.Right)
		} else if right1 != nil {
			cur.Right = right1
		} else if right2 != nil {
			cur.Right = right2
		}
	}
	return merged
}

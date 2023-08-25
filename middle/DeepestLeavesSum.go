package middle

// Definition for a binary tree node.

import "github.com/conglindeng/leetcode/mystruct"

var maxDepth = 0
var res = 0

func DeepestLeavesSum(root *mystruct.TreeNode) int {
	dfs_(root, 0)
	return res
}

func dfs_(root *mystruct.TreeNode, depth int) {
	if root == nil {
		return
	}
	if depth > maxDepth {
		maxDepth = depth
		res = root.Val
	} else if depth == maxDepth {
		res += root.Val
	}
	dfs_(root.Left, depth+1)
	dfs_(root.Right, depth+1)
}

func DeepestLeavesSum_(root *mystruct.TreeNode) int {
	queue := make([]*mystruct.TreeNode, 0)
	res := 0
	max := 0
	depth := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := make([]*mystruct.TreeNode, 0)
		for len(queue) > 0 {
			poll := queue[0]
			queue = queue[1:]
			if depth > max {
				res = poll.Val
				max = depth
			} else if depth == max {
				res += poll.Val
			}

			if poll.Left != nil {
				temp = append(temp, poll.Left)
			}
			if poll.Right != nil {
				temp = append(temp, poll.Right)
			}
		}
		queue = append(queue, temp...)
		depth++
	}
	return res
}

package middle

import (
	"container/list"

	"github.com/conglindeng/leetcode/mystruct"
)

func AddOneRow(root *mystruct.TreeNode, val int, depth int) *mystruct.TreeNode {
	// dfs
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &mystruct.TreeNode{Val: val, Left: root}
	} else if depth == 2 {
		root.Left = &mystruct.TreeNode{Val: val, Left: root.Left}
		root.Right = &mystruct.TreeNode{Val: val, Right: root.Right}
		return root
	} else {
		root.Left = AddOneRow(root.Left, val, depth-1)
		root.Right = AddOneRow(root.Right, val, depth-1)
		return root
	}
}

func AddOneRow_bfs(root *mystruct.TreeNode, val int, depth int) *mystruct.TreeNode {
	mockRoot := &mystruct.TreeNode{Left: root, Val: -1}
	work := list.New()
	work.PushBack(mockRoot)
	level := 0
	for work.Len() != 0 {
		temp := list.New()
		for work.Len() != 0 {
			front := work.Front()
			node := front.Value.(*mystruct.TreeNode)
			if level == depth-1 {
				node.Left = &mystruct.TreeNode{Val: val, Left: node.Left}
				node.Right = &mystruct.TreeNode{Val: val, Right: node.Right}
			} else {
				if node.Left != nil {
					temp.PushBack(node.Left)
				}
				if node.Right != nil {
					temp.PushBack(node.Right)
				}
			}
			work.Remove(front)
		}
		level++
		work.PushBackList(temp)
	}
	return mockRoot.Left
}

package middle

import "github.com/conglindeng/leetcode/mystruct"

func insertIntoMaxTree(root *mystruct.TreeNode, val int) *mystruct.TreeNode {
	if root == nil || root.Val < val {
		root = &mystruct.TreeNode{Val: val, Left: root}
	} else {
		root.Right = insertIntoMaxTree(root.Right, val)
	}
	return root
}

func insertIntoMaxTree_(root *mystruct.TreeNode, val int) *mystruct.TreeNode {
	node := &mystruct.TreeNode{Val: val}
	if root == nil {
		return node
	}
	if root.Val < val {
		node.Left = root
		return node
	}
	temp := root.Right
	pre := root
	for temp != nil && temp.Val > val {
		pre = temp
		temp = temp.Right
	}
	if temp == nil {
		pre.Right = node
	} else {
		pre.Right = node
		node.Left = temp
	}
	return root
}

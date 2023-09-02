package middle

import "github.com/conglindeng/leetcode/mystruct"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func CountNodes(root *mystruct.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	height := 0
	tmp := root
	for tmp != nil {
		height++
		tmp = tmp.Left
	}
	mockRootTree := &mystruct.TreeNode{Left: root}
	find := func(root *mystruct.TreeNode, position, height int) bool {
		for i := 0; i < height; i++ {
			m := (1 << (height - i - 1)) & position
			if m == 0 {
				root = root.Left
			} else {
				root = root.Right
			}
		}
		return root != nil
	}
	left, right := 0, (1 << (height - 1))
	//find last one is false
	for left < right {
		mid := (left + right) / 2
		if find(mockRootTree, mid, height) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if height > 1 {
		res := 0
		for i := 0; i < height-1; i++ {
			res += 1 << i
		}
		return res + left
	}
	return left
}

func countNodes_iterate(root *TreeNode) int {
	var iterate func(root *TreeNode) int
	iterate = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + iterate(root.Left) + iterate(root.Right)
	}
	return iterate(root)
}

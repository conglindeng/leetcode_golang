package middle

import (
	"github.com/conglindeng/leetcode/mystruct"
)

var (
	minWidth map[int]int
	maxWidth int
)

func WidthOfBinaryTree(root *mystruct.TreeNode) int {
	maxWidth = 0
	minWidth = map[int]int{}
	dfs_width(root, 1, 1)
	return maxWidth
}

func dfs_width(root *mystruct.TreeNode, level, width int) {
	if root == nil {
		return
	}
	if _, ok := minWidth[level]; !ok {
		minWidth[level] = width
	}
	maxWidth = max(maxWidth, width-minWidth[level]+1)
	if root.Left != nil {
		dfs_width(root.Left, level+1, 2*width)
	}
	if root.Right != nil {
		dfs_width(root.Right, level+1, 2*width+1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func widthOfBinaryTree(root *mystruct.TreeNode) int {
	min := map[int]int{}
	var dfs func(*mystruct.TreeNode, int, int) int
	dfs = func(root *mystruct.TreeNode, level, width int) int {
		if root == nil {
			return 0
		}
		if _, ok := min[level]; !ok {
			min[level] = width
		}
		return max(width-min[level]+1, max(dfs(root.Left, level+1, width*2), dfs(root.Right, level+1, width*2+1)))
	}
	return dfs(root, 1, 1)
}

type indexAndNodePair struct {
	node  *mystruct.TreeNode
	index int
}

// bfs
func widthOfBinaryTree_(root *mystruct.TreeNode) int {
	res := 0
	q := []indexAndNodePair{{root, 1}}
	for q != nil {
		res = max(res, q[len(q)-1].index-q[0].index+1)
		temp := q
		q = nil
		for _, pair := range temp {
			if pair.node.Left != nil {
				q = append(q, indexAndNodePair{pair.node.Left, pair.index * 2})
			}
			if pair.node.Right != nil {
				q = append(q, indexAndNodePair{pair.node.Right, pair.index*2 + 1})
			}
		}
	}
	return res
}

package middle

import "github.com/conglindeng/leetcode/mystruct"

var (
	maxPathLength int
)

func LongestUnivaluePath(root *mystruct.TreeNode) int {
	if root == nil {
		return 0
	}
	maxPathLength = 0
	dfs_Univalue_Path(root)
	return maxPathLength
}

func dfs_Univalue_Path(root *mystruct.TreeNode) (value, length int) {
	value, length = root.Val, 0
	endLength := 1
	if root.Left != nil {
		leftV, leftL := dfs_Univalue_Path(root.Left)
		if leftV == value {
			length = leftL
			endLength += leftL
		}
	}
	if root.Right != nil {
		rightV, rightL := dfs_Univalue_Path(root.Right)
		if rightV == value {
			length = max(rightL, length)
			endLength += rightL
		}
	}
	maxPathLength = max(maxPathLength, endLength)
	return value, length + 1
}

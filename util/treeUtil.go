package util

import (
	"math"

	"github.com/conglindeng/leetcode/mystruct"
)

func GetHeight(root *mystruct.TreeNode) int {
	if root == nil {
		return 0
	}
	leftH := GetHeight(root.Left)
	rightH := GetHeight(root.Right)
	return int(math.Max(float64(leftH), float64(rightH))) + 1
}

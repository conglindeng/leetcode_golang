package middle

/**
 * Definition for a binary tree node.
 * type mystruct.TreeNode struct {
 *     Val int
 *     Left *mystruct.TreeNode
 *     Right *mystruct.TreeNode
 * }
 */

import "github.com/conglindeng/leetcode/mystruct"

func MaxAncestorDiff(root *mystruct.TreeNode) int {
	type DiffInfo struct {
		Max, Min, Res int
	}
	var dfs func(root *mystruct.TreeNode) DiffInfo
	dfs = func(root *mystruct.TreeNode) DiffInfo {
		if root == nil {
			return DiffInfo{-1, -1, 0}
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		mini, maxi := root.Val, root.Val
		if left.Min != -1 {
			mini = min(left.Min, mini)
			maxi = max(left.Min, maxi)
		}
		if left.Max != -1 {
			mini = min(left.Max, mini)
			maxi = max(left.Max, maxi)
		}
		if right.Max != -1 {
			mini = min(right.Max, mini)
			maxi = max(right.Max, maxi)
		}
		if right.Min != -1 {
			mini = min(right.Min, mini)
			maxi = max(right.Min, maxi)
		}
		res := max(left.Res, right.Res)
		cur := max(root.Val-mini, maxi-root.Val)
		return DiffInfo{maxi, mini, max(res, cur)}
	}
	return dfs(root).Res
}

func MaxAncestorDiff_(root *mystruct.TreeNode) int {
	var dfs func(mini, maxi int, root *mystruct.TreeNode) int
	dfs = func(mini, maxi int, root *mystruct.TreeNode) int {
		if root == nil {
			return 0
		}
		diff := max(abs(root.Val-maxi), abs(root.Val-mini))
		maxi, mini = max(maxi, root.Val), min(mini, root.Val)
		diff = max(diff, dfs(mini, maxi, root.Left))
		diff = max(diff, dfs(mini, maxi, root.Right))
		return diff
	}
	return dfs(root.Val, root.Val, root)
}

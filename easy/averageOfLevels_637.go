package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	sum := map[int]float64{}
	count := map[int]int{}
	var dfs func(root *TreeNode, h int)
	maxH := -1
	dfs = func(root *TreeNode, h int) {
		if root == nil {
			return
		}
		maxH = max(maxH, h)
		sum[h] += float64(root.Val)
		count[h]++
		dfs(root.Left, h+1)
		dfs(root.Right, h+1)
	}
	dfs(root, 0)
	res := []float64{}
	for i := 0; i <= maxH; i++ {
		res = append(res, sum[i]/float64(count[i]))
	}
	return res
}

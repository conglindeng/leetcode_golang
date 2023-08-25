package middle

import (
	"fmt"

	"github.com/conglindeng/leetcode/mystruct"
)

func findDuplicateSubtrees(root *mystruct.TreeNode) []*mystruct.TreeNode {
	repeat := map[*mystruct.TreeNode]struct{}{}
	seen := map[string]*mystruct.TreeNode{}
	var dfs func(*mystruct.TreeNode) string
	dfs = func(root *mystruct.TreeNode) string {
		if root == nil {
			return ""
		}
		str := fmt.Sprintf("%d((%s)(%s))", root.Val, dfs(root.Left), dfs(root.Right))
		fmt.Println(str)
		if n, ok := seen[str]; ok {
			repeat[n] = struct{}{}
		} else {
			seen[str] = root
		}
		return str
	}
	dfs(root)
	res := make([]*mystruct.TreeNode, 0, len(repeat))
	for item := range repeat {
		res = append(res, item)
	}
	return res
}

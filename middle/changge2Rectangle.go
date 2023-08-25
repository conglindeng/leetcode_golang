package middle

import (
	"strconv"

	"github.com/conglindeng/leetcode/mystruct"
	"github.com/conglindeng/leetcode/util"
)

func printTree(root *mystruct.TreeNode) [][]string {
	h := util.GetHeight(root)
	w := (1 << (h + 1)) - 1
	res := make([][]string, h+1)
	for i := 0; i <= h; i++ {
		res[i] = make([]string, w)
	}
	doPrintTree(root, &res, 0, (w-1)/2)
	return res
}

func doPrintTree(root *mystruct.TreeNode, res *[][]string, i, j int) {
	if root != nil {
		(*res)[i][j] = strconv.Itoa(root.Val)
	}
	h := len((*res)) - 1
	if root.Left != nil {
		doPrintTree(root.Left, res, i+1, j-(1<<(h-i-1)))
	}
	if root.Right != nil {
		doPrintTree(root.Right, res, i+1, j+(1<<(h-i-1)))
	}
}

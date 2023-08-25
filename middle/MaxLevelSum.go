package middle

import (
	"math"

	"github.com/conglindeng/leetcode/mystruct"
)

func maxLevelSum(root *mystruct.TreeNode) int {
	sums := make([]int, 0)
	sums = append(sums, math.MaxInt)
	dfs(1, root, &sums)
	res, max := 1, math.MinInt
	for index, item := range sums {
		if item > max && index > 0 {
			res = index
			max = item
		}
	}
	return res
}

func dfs(level int, root *mystruct.TreeNode, sums *[]int) {
	for len(*sums) <= level {
		*sums = append(*sums, 0)
	}
	(*sums)[level] += root.Val
	if root.Left != nil {
		dfs(level+1, root.Left, sums)
	}
	if root.Right != nil {
		dfs(level+1, root.Right, sums)
	}
}

func main() {

}

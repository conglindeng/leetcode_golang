package middle

//   Definition for a binary tree node.

/*   type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 } */
import "github.com/conglindeng/leetcode/mystruct"

// [4,1,-1,2,-1,3]
func Rob(root *mystruct.TreeNode) int {
	cache := map[int][]int{}
	var calcu func(root *mystruct.TreeNode, prePick bool, i int) int
	calcu = func(root *mystruct.TreeNode, prePick bool, i int) int {
		if root == nil {
			cache[i] = []int{0, 0}
			return 0
		}
		var nonPick, pick int
		if _, b := cache[i]; b {
			nonPick = cache[i][0]
			pick = cache[i][1]
		} else {
			nonPick = calcu(root.Left, false, i*2+1) + calcu(root.Right, false, i*2+2)
			pick = root.Val + calcu(root.Left, true, i*2+1) + calcu(root.Right, true, i*2+2)
			cache[i] = []int{nonPick, pick}
		}
		if prePick {
			return nonPick
		} else {
			return max(nonPick, pick)
		}
	}
	return max(calcu(root, true, 0), calcu(root, false, 0))
}

func Rob_(root *mystruct.TreeNode) int {
	var dfs func(root *mystruct.TreeNode) []int
	dfs = func(root *mystruct.TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		pick := root.Val + l[1] + r[1]
		nonPick := max(l[0], l[1]) + max(r[0], r[1])
		return []int{pick, nonPick}
	}
	res := dfs(root)
	return max(res[0], res[1])
}

//4116  [79,99,77,-1,-1,-1,69,-1,60,53,-1,73,11,-1,-1,-1,62,27,62,-1,-1,98,50,-1,-1,90,48,82,-1,-1,-1,55,64,-1,-1,73,56,6,47,-1,93,-1,-1,75,44,30,82,-1,-1,-1,-1,-1,-1,57,36,89,42,-1,-1,76,10,-1,-1,-1,-1,-1,32,4,18,-1,-1,1,7,-1,-1,42,64,-1,-1,39,76,-1,-1,6,-1,66,8,96,91,38,38,-1,-1,-1,-1,74,42,-1,-1,-1,10,40,5,-1,-1,-1,-1,28,8,24,47,-1,-1,-1,17,36,50,19,63,33,89,-1,-1,-1,-1,-1,-1,-1,-1,94,72,-1,-1,79,25,-1,-1,51,-1,70,84,43,-1,64,35,-1,-1,-1,-1,40,78,-1,-1,35,42,98,96,-1,-1,82,26,-1,-1,-1,-1,48,91,-1,-1,35,93,86,42,-1,-1,-1,-1,0,61,-1,-1,67,-1,53,48,-1,-1,82,30,-1,97,-1,-1,-1,1,-1,-1]

package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	stack := []int{}
	stack = append(stack, root.Val)
	res := 1
	var dfs func(root *TreeNode, stack []int)
	dfs = func(root *TreeNode, stack []int) {
		if root == nil {
			return
		}
		if stack[len(stack)-1] <= root.Val {
			res++
			stack = append(stack, root.Val)
		} else {
			stack = append(stack, stack[len(stack)-1])
		}
		dfs(root.Left, stack)
		dfs(root.Right, stack)
		stack = stack[:len(stack)-1]
	}
	dfs(root.Left, stack)
	dfs(root.Right, stack)
	return res
}

func goodNodes_dfs_again(root *TreeNode) int {
	var dfs func(root *TreeNode, pathMax int) int
	dfs = func(root *TreeNode, pathMax int) int {
		if root == nil {
			return 0
		}
		res := 0
		if pathMax <= root.Val {
			res++
			pathMax = root.Val
		}
		res += dfs(root.Left, pathMax) + dfs(root.Right, pathMax)
		return res
	}
	return dfs(root, root.Val)
}

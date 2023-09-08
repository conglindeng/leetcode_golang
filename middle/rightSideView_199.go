package middle

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	work := []*TreeNode{}
	work = append(work, root)
	res := []int{}
	for len(work) > 0 {
		nextLevel := []*TreeNode{}
		for len(work) > 0 {
			cur := work[0]
			work = work[1:]
			if len(work) == 0 {
				res = append(res, cur.Val)
			}
			if cur.Left != nil {
				nextLevel = append(nextLevel, cur.Left)
			}
			if cur.Right != nil {
				nextLevel = append(nextLevel, cur.Right)
			}
		}
		work = nextLevel
	}
	return res
}

func rightSideView_dfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodeStack := []*TreeNode{root}
	heightStack := []int{0}
	lastOne := map[int]int{}
	maxHeight := -1
	for len(nodeStack) > 0 {
		curNode := nodeStack[0]
		nodeStack = nodeStack[1:]
		curHeight := heightStack[0]
		heightStack = heightStack[1:]
		lastOne[curHeight] = curNode.Val
		maxHeight = max(maxHeight, curHeight)
		if curNode.Left != nil {
			nodeStack = append(nodeStack, curNode.Left)
			heightStack = append(heightStack, curHeight+1)
		}
		if curNode.Right != nil {
			nodeStack = append(nodeStack, curNode.Right)
			heightStack = append(heightStack, curHeight+1)
		}
	}
	res := []int{}
	for i := 0; i <= maxHeight; i++ {
		res = append(res, lastOne[i])
	}
	return res
}

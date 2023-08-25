package middle

import "github.com/conglindeng/leetcode/mystruct"

func DelNodes(root *mystruct.TreeNode, to_delete []int) []*mystruct.TreeNode {
	var iterate func(root, pre *mystruct.TreeNode, index int, isForest bool)
	res := []*mystruct.TreeNode{}
	iterate = func(root, pre *mystruct.TreeNode, index int, isForest bool) {
		if root == nil {
			return
		}
		if containNum(to_delete, root.Val) {
			//disconnect with parent
			if !isForest {
				if index%2 == 0 {
					pre.Left = nil
				} else {
					pre.Right = nil
				}
			}
			//disconnect with childrens
			left := root.Left
			right := root.Right
			root.Left = nil
			root.Right = nil
			iterate(left, nil, index*2, true)
			iterate(right, nil, index*2+1, true)
		} else {
			if isForest {
				res = append(res, root)
			}
			iterate(root.Left, root, index*2, false)
			iterate(root.Right, root, index*2+1, false)
		}
	}
	iterate(root, nil, 1, true)
	return res
}

func containNum(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

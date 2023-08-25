package middle

import "github.com/conglindeng/leetcode/mystruct"

func trimBST(root *mystruct.TreeNode, low int, high int) *mystruct.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func trimBST_(root *mystruct.TreeNode, low int, high int) *mystruct.TreeNode {
	for root != nil && (root.Val > high || root.Val < low) {
		if root.Val > high {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	if root == nil {
		return nil
	}
	temp := root
	for temp.Left != nil {
		if temp.Left.Val < low {
			temp.Left = temp.Left.Right
		} else {
			temp = temp.Left
		}
	}
	temp = root
	for temp.Right != nil {
		if temp.Right.Val > high {
			temp.Right = temp.Right.Left
		} else {
			temp = temp.Right
		}
	}
	return root
}

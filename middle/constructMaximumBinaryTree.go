package middle

import (
	"github.com/conglindeng/leetcode/mystruct"
)

func constructMaximumBinaryTree(nums []int) *mystruct.TreeNode {

	return doBuild(nums, 0, len(nums)-1)
}

func doBuild(nums []int, start, end int) *mystruct.TreeNode {
	if end < start {
		return nil
	}
	if start == end {
		return &mystruct.TreeNode{Val: nums[start]}
	}
	index := findMax(nums, start, end)
	res := &mystruct.TreeNode{Val: nums[index]}
	res.Left = doBuild(nums, start, index-1)
	res.Right = doBuild(nums, index+1, end)
	return res
}

func findMax(nums []int, start, end int) int {
	res := start
	for i := start; i <= end; i++ {
		if nums[i] > nums[res] {
			res = i
		}
	}
	return res
}

func ConstructMaximumBinaryTree_(nums []int) *mystruct.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	best := 0
	for i, num := range nums {
		if num > nums[best] {
			best = i
		}
	}
	return &mystruct.TreeNode{Val: nums[best], Left: ConstructMaximumBinaryTree_(nums[:best]), Right: ConstructMaximumBinaryTree_(nums[best+1:])}
}

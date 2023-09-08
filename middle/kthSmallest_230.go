package middle

import (
	"container/heap"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	h := &KthHeap{}
	heap.Init(h)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		heap.Push(h, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	for k > 1 {
		heap.Pop(h)
		k--
	}
	return heap.Pop(h).(int)
}

type KthHeap struct {
	sort.IntSlice
}

func (h *KthHeap) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *KthHeap) Pop() any {
	cur := h.IntSlice[len(h.IntSlice)]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return cur
}

func kthSmallest_LDR(root *TreeNode, k int) int {
	var ldr func(root *TreeNode) []int
	ldr = func(root *TreeNode) []int {
		work := []*TreeNode{}
		cur := root
		res := []int{}
		for cur != nil || len(work) > 0 {
			if cur != nil {
				work = append(work, cur)
				cur = cur.Left
			} else {
				cur = work[len(work)-1]
				work = work[:len(work)-1]
				res = append(res, cur.Val)
				cur = cur.Right
			}
		}
		return res
	}
	nums := ldr(root)
	return nums[k-1]
}

func kthSmallest_dfs(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	cur := root
	for {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		stack = stack[:len(stack)-1]
		cur = stack[len(stack)-1]
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
}

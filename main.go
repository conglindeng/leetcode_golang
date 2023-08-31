package main

import (
	"fmt"

	"github.com/conglindeng/leetcode/converter"
	_ "github.com/conglindeng/leetcode/converter"
	"github.com/conglindeng/leetcode/difficult"
	"github.com/conglindeng/leetcode/easy"
	"github.com/conglindeng/leetcode/middle"
	"github.com/conglindeng/leetcode/mystruct"
)

func main() {
	listNodeMain()
	t := converter.ConvertIntArray2Tree([]int{8, 3, 10, 1, 6, -1, 14, -1, -1, 4, 7, 13}, -1)
	t = &mystruct.TreeNode{Val: 1}
	t.Left = &mystruct.TreeNode{Val: 5}
	t.Left.Left = &mystruct.TreeNode{Val: 8}
	t.Left.Left.Right = &mystruct.TreeNode{Val: 10}
	t.Right = &mystruct.TreeNode{Val: 2}
	t.Right.Right = &mystruct.TreeNode{Val: 0}
	t.Right.Right.Left = &mystruct.TreeNode{Val: 3}

	easy.ÇircularGameLosers(4, 4)

	middle.NumFactoredBinaryTrees([]int{2, 4, 5, 10})
	// r := middle.MinSubarray([]int{6, 3, 5, 2}, 9)

}

func listNodeMain() {

	list1 := mystruct.ChangeArray2ListNode([]int{9, 5, 6, 4, 7, 13}, 2)
	middle.DetectCycle(list1)

	first := &difficult.ListNode{Val: 1}
	first.Next = &difficult.ListNode{Val: 4}
	first.Next.Next = &difficult.ListNode{Val: 5}

	second := &difficult.ListNode{Val: 1}
	second.Next = &difficult.ListNode{Val: 3}
	second.Next.Next = &difficult.ListNode{Val: 4}

	third := &difficult.ListNode{Val: 2}
	third.Next = &difficult.ListNode{Val: 6}

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums[:2])
	fmt.Println(nums[2:])

	ln := difficult.MergeKLists([]*difficult.ListNode{first, second, third})
	for ln != nil {
		fmt.Printf(" -> %d ", ln.Val)
		ln = ln.Next
	}
}

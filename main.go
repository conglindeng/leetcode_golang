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
	t := converter.ConvertIntArray2Tree([]int{1, 2, 3, 4, 5, 6}, -1)
	fmt.Println(t)
	t2 := &mystruct.TreeNode{Val: 1}
	t2.Left = &mystruct.TreeNode{Val: 5}
	t2.Left.Left = &mystruct.TreeNode{Val: 8}
	t2.Left.Left.Right = &mystruct.TreeNode{Val: 10}
	t2.Right = &mystruct.TreeNode{Val: 2}
	t2.Right.Right = &mystruct.TreeNode{Val: 0}
	t2.Right.Right.Left = &mystruct.TreeNode{Val: 3}
	middle.CountNodes(t)

	easy.CaptureForts([]int{1, 0, 0, -1, 0, 0, 0, 1, 0, -1})

	middle.EvalRPN_forward([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	// r := middle.MinSubarray([]int{6, 3, 5, 2}, 9)

	difficult.Calculate("- 2+1")

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

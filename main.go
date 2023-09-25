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
	t := converter.ConvertIntArray2Tree([]int{79, 99, 77, -1, -1, -1, 69, -1, 60, 53, -1, 73, 11, -1, -1, -1, 62, 27, 62, -1, -1, 98, 50, -1, -1, 90, 48, 82, -1, -1, -1, 55, 64, -1, -1, 73, 56, 6, 47, -1, 93, -1, -1, 75, 44, 30, 82, -1, -1, -1, -1, -1, -1, 57, 36, 89, 42, -1, -1, 76, 10, -1, -1, -1, -1, -1, 32, 4, 18, -1, -1, 1, 7, -1, -1, 42, 64, -1, -1, 39, 76, -1, -1, 6, -1, 66, 8, 96, 91, 38, 38, -1, -1, -1, -1, 74, 42, -1, -1, -1, 10, 40, 5, -1, -1, -1, -1, 28, 8, 24, 47, -1, -1, -1, 17, 36, 50, 19, 63, 33, 89, -1, -1, -1, -1, -1, -1, -1, -1, 94, 72, -1, -1, 79, 25, -1, -1, 51, -1, 70, 84, 43, -1, 64, 35, -1, -1, -1, -1, 40, 78, -1, -1, 35, 42, 98, 96, -1, -1, 82, 26, -1, -1, -1, -1, 48, 91, -1, -1, 35, 93, 86, 42, -1, -1, -1, -1, 0, 61, -1, -1, 67, -1, 53, 48, -1, -1, 82, 30, -1, 97, -1, -1, -1, 1, -1, -1}, -1)
	fmt.Println(t)
	// t2 := &mystruct.TreeNode{Val: 1}
	// t2.Left = &mystruct.TreeNode{Val: 5}
	// t2.Left.Left = &mystruct.TreeNode{Val: 8}
	// t2.Left.Left.Right = &mystruct.TreeNode{Val: 10}
	// t2.Right = &mystruct.TreeNode{Val: 2}
	// t2.Right.Right = &mystruct.TreeNode{Val: 0}
	// t2.Right.Right.Left = &mystruct.TreeNode{Val: 3}
	middle.Rob_tree_again(t)

	easy.CaptureForts([]int{1, 0, 0, -1, 0, 0, 0, 1, 0, -1})

	middle.Rob__([]int{1})
	// r := middle.MinSubarray([]int{6, 3, 5, 2}, 9)

	difficult.Calculate("- 2+1")

	l := difficult.Constructor(2)
	l.Put(3, 1)
	l.Put(2, 1)// 2,1  1,1
	l.Put(2, 2)// 2,3  1,1
	l.Put(4, 4) //2,3  4,1
	l.Get(2)
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

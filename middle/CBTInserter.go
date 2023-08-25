package middle

import (
	"math/bits"

	"github.com/conglindeng/leetcode/mystruct"
)

type CBTInserter struct {
	root  *mystruct.TreeNode
	count int
}

func CBTInserter_Constructor(root *mystruct.TreeNode) CBTInserter {
	work := []*mystruct.TreeNode{root}
	count := 0
	for len(work) > 0 {
		count++
		poll := work[0]
		work = work[1:]
		if poll.Left != nil {
			work = append(work, poll.Left)
		}
		if poll.Right != nil {
			work = append(work, poll.Right)
		}
	}
	return CBTInserter{count: count, root: root}
}

func (this *CBTInserter) Insert(val int) int {
	this.count++
	child := &mystruct.TreeNode{Val: val}
	temp := this.root
	for i := bits.Len(uint(this.count)) - 2; i > 0; i-- {
		if (this.count>>i)&1 == 1 {
			temp = temp.Right
		} else {
			temp = temp.Left
		}
	}
	if this.count&1 == 1 {
		temp.Right = child
	} else {
		temp.Left = child
	}
	return temp.Val
}

func (this *CBTInserter) Get_root() *mystruct.TreeNode {
	return this.root
}

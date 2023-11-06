package middle

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	work := []*Node{}
	work = append(work, root)
	for len(work) > 0 {
		tmp := []*Node{}
		for len(work) > 0 {
			first := work[0]
			work = work[1:]
			if len(work) > 0 {
				first.Next = work[0]
			}
			if first.Left != nil {
				tmp = append(tmp, first.Left)
			}
			if first.Right != nil {
				tmp = append(tmp, first.Right)
			}
		}
		work = append(work, tmp...)
	}
	return root
}

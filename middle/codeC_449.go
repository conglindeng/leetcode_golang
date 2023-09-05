package middle

import (
	"math"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor_CodeC() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	vals := []string{}
	var lrd func(root *TreeNode)
	lrd = func(root *TreeNode) {
		if root != nil {
			lrd(root.Left)
			lrd(root.Right)
			vals = append(vals, strconv.Itoa(root.Val))
		}
	}
	lrd(root)
	return strings.Join(vals, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, ",")
	var build func(min, max int) *TreeNode
	build = func(min, max int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < min || val > max {
			return nil
		}
		arr = arr[:len(arr)-1]
		return &TreeNode{Val: val, Right: build(val, max), Left: build(min, val)}
	}
	return build(math.MinInt, math.MaxInt)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

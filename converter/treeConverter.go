package converter

import (
	"strconv"

	"github.com/conglindeng/leetcode/mystruct"
)

func ConvertIntArray2Tree(values []int, nullNum int) *mystruct.TreeNode {
	strs := make([]string, 0)
	for _, num := range values {
		if num == nullNum {
			strs = append(strs, "")
		} else {
			strs = append(strs, strconv.Itoa(num))
		}
	}
	return ConvertStrArray2Tree(strs)
}

func ConvertStrArray2Tree(values []string) *mystruct.TreeNode {
	return doConvert2Tree(values, 0)
}

func doConvert2Tree(values []string, index int) *mystruct.TreeNode {
	if index < len(values) {
		if values[index] == "" {
			return nil
		}
		n, _ := strconv.Atoi(values[index])
		node := &mystruct.TreeNode{Val: n}
		node.Left = doConvert2Tree(values, 2*index+1)
		node.Right = doConvert2Tree(values, 2*index+2)
		return node
	}
	return nil
}

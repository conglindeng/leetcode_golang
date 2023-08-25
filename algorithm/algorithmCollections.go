package algorithm

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/conglindeng/leetcode/mystruct"
)

// Write a program to check if a binary tree is a valid binary search tree.
func ValidateBinarySearchTree(root *mystruct.TreeNode) bool {
	// l-d-r traverse is incremental
	if root == nil {
		return false
	}
	stack := make([]*mystruct.TreeNode, 0)
	temp := root
	pre := math.MinInt
	for len(stack) > 0 || temp != nil {
		if temp != nil {
			stack = append(stack, temp)
			temp = temp.Left
		} else {
			temp = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if temp.Val < pre {
				return false
			}
			pre = temp.Val
			temp = temp.Right
		}
	}
	return true
}

// For a given rotated array that is sorted, write a program code to find a particular element in the array.
func find(nums []int) int {
	left, right := 0, len(nums)-1
	var mid int
	mid = 0
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] < 0 {

		}
	}
	return -1
}

// For a given rotated array that is sorted, write a program code to find a particular element in the array.
func FindInRotatedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid+1] <= nums[right] {
			if nums[mid+1] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target <= nums[mid-1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

// Write an error-free program to check if a given binary tree is balanced or not.
func ValidateTreeIsBalance(root *mystruct.TreeNode) bool {
	_, b := doValidateTreeIsBalance(root, 0)
	return b
}

func doValidateTreeIsBalance(root *mystruct.TreeNode, height int) (int, bool) {
	if root == nil {
		return height, true
	}
	leftH, leftB := doValidateTreeIsBalance(root.Left, height)
	rightH, rightB := doValidateTreeIsBalance(root.Right, height)
	h := int(math.Max(float64(leftH), float64(rightH)))
	m := int(math.Abs(float64(leftH) - float64(rightH)))
	return h + 1, m <= 1 && leftB && rightB

}

// You are provided with a string in the form of a sentence. Write a program to display the characters of the string in reverse order.
func rotateSentence(sentence string) string {
	s := strings.Split(sentence, " ")
	var b bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteString(s[i])
		b.WriteString(" ")
	}
	return b.String()
}

// [1,3,4]
// For a given unsorted array with positive integers from 1 to n and one missing element, write a program to find the missing element in the unsorted array.
func findMissOne(nums []int) int {
	l := len(nums)
	sum := l * (1 + l) / 2
	for _, n := range nums {
		sum -= n
	}
	return sum
}

// move zeros to the end of array. example [1,0,3,0,1] -> [1,3,1,0,0]
func MoveZeros(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for ; index < len(nums); index++ {
		nums[index] = 0
	}
	fmt.Print(nums)
}

func FindMissTwo(nums []int) []int {
	sort.Ints(nums)
	l := len(nums) + 2
	res := make([]int, 0)
	m := 1
	for i := 0; i < l-2; {
		if nums[i] != m {
			res = append(res, m)
		} else {
			i++
		}
		m++
	}
	for m <= l {
		res = append(res, m)
		m++
	}
	return res
}

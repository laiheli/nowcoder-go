package problems

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	var head, ptr *ListNode
	for _, v := range nums {
		if ptr == nil {
			ptr = &ListNode{v, nil}
			head = ptr
		} else {
			ptr.Next = &ListNode{v, nil}
			ptr = ptr.Next
		}
	}
	return head

}

func (l *ListNode) String() string {
	var data []string
	var ptr = l

	for ptr != nil {
		ptr, data = ptr.Next, append(data, strconv.Itoa(ptr.Val))
	}

	return strings.Join(data, `,`)
}

func (l *ListNode) Equal(list *ListNode) bool {
	var l1, l2 = l, list
	for l1 != nil {
		if l2 == nil || l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	return l2 == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

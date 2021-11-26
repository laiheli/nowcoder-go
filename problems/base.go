package problems

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode 切片转链表
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

// NewCycleListNode 切片转带环链表
func NewCycleListNode(nums []int, pos int) *ListNode {
	var head, ptr, pPtr *ListNode

	for k, v := range nums {
		if ptr == nil {
			ptr = &ListNode{v, nil}
			head = ptr
		} else {
			ptr.Next = &ListNode{v, nil}
			ptr = ptr.Next
		}
		if k == pos-1 {
			pPtr = ptr
		}
	}

	if ptr != nil {
		ptr.Next = pPtr
	}

	return head

}

func (l *ListNode) String() string {
	var inCycle, entryNode = false, entryNodeOfLoop(l)
	var res []string
	var tmp string
	var ptr = l

	for ptr != nil {
		tmp = ``
		// 有环
		if ptr == entryNode {
			// 第二次遇到环入口，链表结束
			if inCycle {
				break
			} else {
				tmp = `->`
				inCycle = true
			}
		}

		ptr, res = ptr.Next, append(res, tmp+strconv.Itoa(ptr.Val))
	}

	return strings.Join(res, `,`)
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

func NewTreeNode(s []int, nilVal int) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	var stack = make([]*TreeNode, len(s))
	var pIndex = 0
	for k, v := range s {
		var node *TreeNode
		if v != nilVal {
			node = &TreeNode{Val: v}
		}
		stack[k] = node
		if k > 0 {
			if k%2 == 1 {
				if stack[pIndex] != nil {
					stack[pIndex].Left = node
				}
			} else {
				if stack[pIndex] != nil {
					stack[pIndex].Right = node
				}
				pIndex++
			}
		}
	}
	return stack[0]
}

func (t *TreeNode) String() string {
	var str = make([]string, 0)
	var appendT func(tn *TreeNode)
	appendT = func(tn *TreeNode) {
		if tn != nil {
			str = append(str, strconv.Itoa(tn.Val))
			appendT(tn.Left)
			appendT(tn.Right)
		}
	}
	appendT(t)
	return strings.Join(str, ",")
}

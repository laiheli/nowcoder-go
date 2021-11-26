package problems

// NC1 大数加法
func solve(s string, t string) string {
	if len(s) > len(t) {
		s, t = t, s
	}

	var res = []byte(t)
	var carry, sum uint8 = 0, 0
	sLen, tLen := len(s), len(t)

	for i := 1; i <= tLen; i++ {
		carry, sum = 0, t[tLen-i]-'0'+carry
		if i <= sLen {
			sum += s[sLen-i] - '0'
		}

		if sum > 9 {
			carry, sum = 1, sum%10
		}

		res[tLen-i] = sum + '0'

		if carry == 0 && i > sLen {
			break
		}
	}

	if carry == 1 {
		return `1` + string(res)
	}

	return string(res)
}

// NC2 重排链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 找到中点
	end, ptr := head, head
	for ptr.Next != nil && ptr.Next.Next != nil {
		end, ptr = end.Next, ptr.Next.Next
	}

	// 倒转链表
	end, end.Next, ptr = end.Next, nil, nil
	for end != nil {
		end, end.Next, ptr = end.Next, ptr, end
	}

	// 合并链表
	end, ptr = ptr, head
	for end != nil {
		ptr, ptr.Next, end, end.Next = ptr.Next, end, end.Next, ptr.Next
	}
}

// NC3 链表中环的入口结点
func entryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}

	var s, f = pHead, pHead
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
		if s == f {
			break
		}
	}

	if f == nil {
		return nil
	}

	f = pHead
	for s != f {
		s, f = s.Next, f.Next
	}
	return s
}

// NC4 判断链表中是否有环
func hasCycle(head *ListNode) bool {
	var s, f = head, head
	for f != nil && f.Next != nil && f.Next.Next != nil {
		s, f = s.Next, f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

// NC5 二叉树根节点到叶子节点的所有路径和
func sumNumbers(root *TreeNode) int {
	var total = 0
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node != nil {
			sum = sum*10 + node.Val
			if node.Left == nil && node.Right == nil {
				total += sum
			}
			if node.Left != nil {
				dfs(node.Left, sum)
			}
			if node.Right != nil {
				dfs(node.Right, sum)
			}
		}
	}
	dfs(root, 0)

	return total
}

// NC6 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	var max = -1001
	var maxNum = func(nums ...int) int {
		if len(nums) == 0 {
			return 0
		}
		m := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > m {
				m = nums[i]
			}
		}
		return m
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node != nil {
			p, cl, cr := node.Val, dfs(node.Left), dfs(node.Right)
			max = maxNum(max, maxNum(p, p+cl, p+cr, p+cl+cr))

			return maxNum(p, p+cl, p+cr)
		}
		return 0
	}

	dfs(root)

	return max
}

// NC7 买卖股票的最好时机(一)
func maxProfit(prices []int) int {
	min, max := 100000, 0
	for _, v := range prices {
		if max < v-min {
			max = v - min
		}
		if min > v {
			min = v
		}
	}

	return max
}

// NC8 二叉树中和为某一值的路径(二)
func findPath(root *TreeNode, expectNumber int) [][]int {
	var res [][]int
	var dfs func(node *TreeNode, nums []int, sum int)
	dfs = func(node *TreeNode, nums []int, sum int) {
		if node != nil {
			if node.Left == nil && node.Right == nil {
				if sum+node.Val == expectNumber {
					res = append(res, append(nums, node.Val))
				}
			} else {

				if node.Left != nil {
					dfs(node.Left, append(nums, node.Val), sum+node.Val)
				}

				if node.Right != nil {
					dfs(node.Right, append(nums, node.Val), sum+node.Val)
				}
			}
		}
	}

	dfs(root, []int{}, 0)

	return res
}

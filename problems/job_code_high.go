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
	var s, f = pHead, pHead
	for {
		s, f = s.Next, f.Next.Next
		if s == f {
			break
		}
	}

	f = pHead
	for s != f {
		s, f = s.Next, f.Next
	}
	return s
}

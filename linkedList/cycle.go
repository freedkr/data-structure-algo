package linkedList

// HasCycle 哈希表
func HasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if ok := m[cur]; ok {
			return ok
		}
		m[cur] = true
		cur = cur.next
	}
	return false
}

// HasCycle2 双指针
func HasCycle2(head *ListNode) bool {
	slow := head
	fast := head.next
	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
}

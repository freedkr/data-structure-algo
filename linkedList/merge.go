package linkedList

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.val > l2.val {
		l1.next = mergeTwoList(l1.next, l2)
		return l1
	} else {
		l2.next = mergeTwoList(l1, l2.next)
		return l2
	}

}

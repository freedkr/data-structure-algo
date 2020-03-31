package linkedList

type ListNode struct {
	val  int
	next *ListNode
}

func Reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.next
		cur.next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

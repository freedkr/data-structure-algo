package linkedList

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := &ListNode{
		val: 3,
		next: &ListNode{
			val: 2,
			next: &ListNode{
				val:  1,
				next: nil,
			},
		},
	}
	n := Reverse(head)
	for n != nil {
		fmt.Println(n)
		n = n.next
	}
}

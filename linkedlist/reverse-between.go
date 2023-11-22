package linkedlist

// ReverseBetween reverses the nodes between indexes (using 0-based indexing) m and n (inclusive) in the linked list.
func (ll *LinkedList) ReverseBetween(m, n int) {
	if ll.Length <= 1 || n <= m {
		return
	}

	if m < 0 {
		m = 0
	}

	if n >= ll.Length-1 {
		n = ll.Length - 1
	}

	dummy := &Node{0, nil}
	dummy.Next = ll.Head
	prev := dummy

	for i := 0; i < m; i++ {
		prev = prev.Next
	}

	curr := prev.Next
	for i := 0; i < n-m; i++ {
		temp := curr.Next
		curr.Next = temp.Next
		temp.Next = prev.Next
		prev.Next = temp
	}

	ll.Head = dummy.Next

	if n == ll.Length-1 {
		ll.Tail = curr
	}
}

// e.g. ReverseBetween(0, 2)
//
// 1->2->3->4->5->nil
//
// p  c  t
// 0->1->2->3->4->5->nil
//
// p  t  c
// 0->2->1->3->4->5->nil
//
// p     c  t
// 0->2->1->3->4->5->nil
//
// p     t  c
// 0->3->2->1->4->5->nil
//
// p     t  c
// 0->3->2->1->4->5->nil
//
// 3->2->1->4->5->nil

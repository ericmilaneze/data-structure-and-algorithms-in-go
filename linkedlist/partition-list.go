package linkedlist

// PartitionList partitions the linked list such that all nodes with values less than x come before nodes with values greater than or equal to x.
//
// Note 1 (from exercise): this linked list class does not have a tail which will make this method easier to implement.
//
// Note 2: The actual list implemented here has a tail (so disregard note 1)
func (ll *LinkedList) PartitionList(x int) {
	if ll.Length <= 1 {
		return
	}

	// creating dummy nodes
	lessThanHead, graterThanOrEqualHead := &Node{}, &Node{}
	var lessThanTail, graterThanOrEqualTail = lessThanHead, graterThanOrEqualHead

	curr := ll.Head
	for curr != nil {
		if curr.Value < x {
			lessThanTail.Next = curr
			lessThanTail = curr
		} else {
			graterThanOrEqualTail.Next = curr
			graterThanOrEqualTail = curr
		}

		curr = curr.Next
	}

	if lessThanHead.Next != nil {
		ll.Head = lessThanHead.Next

		if graterThanOrEqualHead.Next != nil {
			ll.Tail = graterThanOrEqualTail
		}
	} else {
		ll.Head = graterThanOrEqualHead.Next
		ll.Tail = graterThanOrEqualTail
	}

	lessThanTail.Next = graterThanOrEqualHead.Next
	graterThanOrEqualTail.Next = nil
}

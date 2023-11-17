package linkedlist

// FindKthNodeFromEnd finds and returns the kth node from the end of the linked list.
//
// Note: this LinkedList implementation does not have a length member variable.
func (ll LinkedList) FindKthNodeFromEnd(k int) *Node {
	slow, fast := ll.Head, ll.Head

	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}

		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

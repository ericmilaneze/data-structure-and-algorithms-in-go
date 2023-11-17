package linkedlist

func (ll LinkedList) HasLoop() bool {
	slow, fast := ll.Head, ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

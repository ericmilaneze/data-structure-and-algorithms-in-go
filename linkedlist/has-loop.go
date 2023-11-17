package linkedlist

// HasLoop detects if there is a cycle or loop present in the linked list.
//
// The method should utilize Floyd's cycle-finding algorithm, also known as the "tortoise and hare" algorithm, to determine the presence of a loop efficiently.
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

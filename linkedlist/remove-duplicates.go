package linkedlist

// RemoveDuplicates removes all duplicate nodes from the linked list based on their values.
func (ll *LinkedList) RemoveDuplicates() {
	if ll.Length <= 1 {
		return
	}

	values := make(map[int]bool)
	values[ll.Head.Value] = true

	prev := ll.Head
	curr := ll.Head.Next

	for curr != nil {
		temp := curr.Next
		currValue := curr.Value

		if _, found := values[curr.Value]; found {
			prev.Next = temp
			curr.Next = nil
			ll.Length--

			if temp == nil {
				ll.Tail = prev
				return
			}
		} else {
			prev = curr
		}

		curr = temp
		values[currValue] = true
	}
}

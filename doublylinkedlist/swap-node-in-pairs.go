package doublylinkedlist

func (dll *DoublyLinkedList) SwapPairs() {
	if dll.Length <= 1 {
		return
	}

	previousTail := dll.Tail.Previous
	current := dll.Head
	dll.Head = current.Next

	for i := 0; i < dll.Length/2; i++ {
		next := current.Next
		temp := next.Next

		if current.Previous != nil {
			current.Previous.Next = next
		}

		next.Previous = current.Previous
		next.Next = current
		current.Next = temp
		current.Previous = next

		if temp != nil {
			temp.Previous = current
		}

		current = temp
	}

	if dll.Length%2 == 0 {
		dll.Tail = previousTail
	}
}

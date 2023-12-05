package doublylinkedlist

func (dll DoublyLinkedList) IsPalindrome() bool {
	if dll.Length <= 1 {
		return true
	}

	nodeFromHead := dll.Head
	nodeFromTail := dll.Tail

	for i := 0; i < dll.Length/2; i++ {
		if nodeFromHead.Value != nodeFromTail.Value {
			return false
		}

		nodeFromHead = nodeFromHead.Next
		nodeFromTail = nodeFromTail.Previous
	}

	return true
}

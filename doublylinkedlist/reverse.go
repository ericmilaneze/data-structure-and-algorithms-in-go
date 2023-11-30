package doublylinkedlist

func (dll *DoublyLinkedList) Reverse() {
	if dll.Length <= 1 {
		return
	}

	fromHead := dll.Head
	fromTail := dll.Tail

	for i := 0; i < dll.Length/2; i++ {
		temp := fromHead.Value
		fromHead.Value = fromTail.Value
		fromTail.Value = temp

		fromHead = fromHead.Next
		fromTail = fromTail.Previous
	}
}

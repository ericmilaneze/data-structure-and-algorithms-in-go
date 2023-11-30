package doublylinkedlist

func (dll *DoublyLinkedList) SwapFirstAndLast() {
	if dll.Length <= 1 {
		return
	}

	temp := dll.Head.Value
	dll.Head.Value = dll.Tail.Value
	dll.Tail.Value = temp
}

package linkedlist

func (ll *LinkedList) FindMiddleNode() *Node {
	if ll.Length == 0 {
		return nil
	}

	slow, fast := ll.Head, ll.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

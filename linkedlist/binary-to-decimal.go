package linkedlist

// BinaryToDecimal converts this binary number, represented by the linked list, into its decimal equivalent.
func (ll LinkedList) BinaryToDecimal() int {
	sum := 0

	curr := ll.Head

	for curr != nil {
		sum = sum*2 + curr.Value

		curr = curr.Next
	}

	return sum
}

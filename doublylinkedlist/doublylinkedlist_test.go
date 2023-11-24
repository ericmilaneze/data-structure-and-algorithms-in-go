package doublylinkedlist

func ExampleDoublyLinkedList() {
	dll := FromValues(10, 11, 12, 13, 14)
	dll.Print()

	//Output:
	// nil<-10<->11<->12<->13<->14->nil
	// nil<-14<->13<->12<->11<->10->nil
	// Head: 10
	// Tail: 14
	// Length: 5
}

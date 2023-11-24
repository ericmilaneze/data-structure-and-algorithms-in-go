// Package doublylinkedlist provides a basic DoublyLinkedList type and methods to handle it, like Push(), Pop(), etc.
package doublylinkedlist

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func (node Node) String() string {
	return fmt.Sprintf("Value: %v - Next: %v - Previous: %v", node.Value, node.Next, node.Previous)
}

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (ll DoublyLinkedList) String() string {
	return fmt.Sprintf("Head: %v\r\nTail: %v\r\nLength: %v", ll.Head, ll.Tail, ll.Length)
}

func New(value int) *DoublyLinkedList {
	node := &Node{0, nil, nil}
	return &DoublyLinkedList{node, node, 1}
}

func FromValues(values ...int) *DoublyLinkedList {
	dll := &DoublyLinkedList{}

	for _, v := range values {
		dll.Push(v)
	}

	return dll
}

func (dll *DoublyLinkedList) Push(value int) *DoublyLinkedList {
	if dll.Length == 0 {
		dll.Head = &Node{value, nil, nil}
		dll.Tail = dll.Head
	} else {
		node := &Node{Value: value, Previous: dll.Tail}
		dll.Tail.Next = node
		dll.Tail = node
	}

	dll.Length++

	return dll
}

// Print shows the details of the Doubly Linked List on the screen.
func (dll DoublyLinkedList) Print() {
	if dll.Length > 0 {
		fmt.Println(dll.SprintValues())
		fmt.Println(dll.SprintValuesReverse())
		fmt.Println("Head:", dll.Head.Value)
		fmt.Println("Tail:", dll.Tail.Value)
	}

	fmt.Println("Length:", dll.Length)
}

// SprintValues returns the values formatted like nil<-1<->2<->3->nil
func (dll DoublyLinkedList) SprintValues() string {
	ret := []byte{}

	if dll.Length == 0 {
		return "nil"
	}

	ret = fmt.Append(ret, "nil<-")

	curr := dll.Head

	for curr != nil {
		ret = fmt.Append(ret, curr.Value)

		if curr.Next != nil {
			ret = fmt.Append(ret, "<->")
		}

		curr = curr.Next
	}

	ret = fmt.Append(ret, "->nil")

	return string(ret)
}

// SprintValuesReverse returns the values formatted like nil<-3<->2<->1->nil
func (dll DoublyLinkedList) SprintValuesReverse() string {
	ret := []byte{}

	if dll.Length == 0 {
		return "nil"
	}

	ret = fmt.Append(ret, "nil<-")

	curr := dll.Tail

	for curr != nil {
		ret = fmt.Append(ret, curr.Value)

		if curr.Previous != nil {
			ret = fmt.Append(ret, "<->")
		}

		curr = curr.Previous
	}

	ret = fmt.Append(ret, "->nil")

	return string(ret)
}

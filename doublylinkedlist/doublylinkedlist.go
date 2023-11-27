// Package doublylinkedlist provides a basic DoublyLinkedList type and methods to handle it, like Push(), Pop(), etc.
package doublylinkedlist

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func (node Node) String() string {
	var r []byte

	r = fmt.Appendf(r, "Value: %v\n", node.Value)

	if node.Next != nil {
		r = fmt.Appendf(r, "Next: %v\n", node.Next.Value)
	}

	if node.Previous != nil {
		r = fmt.Appendf(r, "Previous: %v\n", node.Previous.Value)
	}

	return string(r)
}

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (dll DoublyLinkedList) String() string {
	var r []byte

	if dll.Head != nil {
		r = fmt.Appendf(r, "Head: %v\n", dll.Head.Value)
	}

	if dll.Tail != nil {
		r = fmt.Appendf(r, "Tail: %v\n", dll.Tail.Value)
	}

	r = fmt.Appendf(r, "Length: %v\n", dll.Length)
	r = fmt.Appendf(r, "Items: %s\n", dll.SprintValues())
	r = fmt.Appendf(r, "Items (reverse): %s", dll.SprintValuesReverse())

	return string(r)
}

func New(value int) *DoublyLinkedList {
	node := &Node{value, nil, nil}
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

func (dll *DoublyLinkedList) Pop() *Node {
	if dll.Length == 0 {
		return nil
	}

	if dll.Length == 1 {
		temp := dll.Tail
		dll.Head = nil
		dll.Tail = nil
		dll.Length = 0

		return temp
	}

	temp := dll.Tail
	dll.Tail = dll.Tail.Previous
	dll.Tail.Previous.Next = nil
	dll.Length--

	return temp
}

func (dll *DoublyLinkedList) Unshift(value int) *DoublyLinkedList {
	node := &Node{value, dll.Head, nil}

	if dll.Length > 0 {
		dll.Head.Previous = node
	} else {
		dll.Tail = node
	}

	dll.Head = node
	dll.Length++

	return dll
}

func (dll *DoublyLinkedList) Shift() *Node {
	if dll.Length == 0 {
		return nil
	}

	if dll.Length == 1 {
		temp := dll.Head
		dll.Head = nil
		dll.Tail = nil
		dll.Length = 0
		return temp
	}

	temp := dll.Head
	dll.Head = dll.Head.Next
	dll.Head.Previous = nil
	dll.Length--

	return temp
}

// SprintValues returns the values formatted like nil<-1<->2<->3->nil
func (dll DoublyLinkedList) SprintValues() string {
	r := []byte{}

	if dll.Length == 0 {
		return "nil"
	}

	r = fmt.Append(r, "nil<-")

	curr := dll.Head

	for curr != nil {
		r = fmt.Append(r, curr.Value)

		if curr.Next != nil {
			r = fmt.Append(r, "<->")
		}

		curr = curr.Next
	}

	r = fmt.Append(r, "->nil")

	return string(r)
}

// SprintValuesReverse returns the values formatted like nil<-3<->2<->1->nil
func (dll DoublyLinkedList) SprintValuesReverse() string {
	r := []byte{}

	if dll.Length == 0 {
		return "nil"
	}

	r = fmt.Append(r, "nil<-")

	curr := dll.Tail
	for curr != nil {
		r = fmt.Append(r, curr.Value)

		if curr.Previous != nil {
			r = fmt.Append(r, "<->")
		}

		curr = curr.Previous
	}

	r = fmt.Append(r, "->nil")

	return string(r)
}

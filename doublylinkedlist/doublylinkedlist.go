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
	dll.Tail.Previous.Next = nil
	dll.Tail = dll.Tail.Previous
	dll.Tail.Next = nil
	temp.Next = nil
	temp.Previous = nil
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

func (dll DoublyLinkedList) Get(index int) *Node {
	if index < 0 || index > dll.Length-1 || dll.Length == 0 {
		return nil
	}

	if index == 0 {
		return dll.Head
	}

	if index == dll.Length-1 {
		return dll.Tail
	}

	if index > dll.Length/2 {
		curr := dll.Tail
		for i := 0; i < dll.Length-1-index; i++ {
			curr = curr.Previous
		}
	}

	curr := dll.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	return curr
}

func (dll *DoublyLinkedList) Set(index, value int) bool {
	node := dll.Get(index)

	if node == nil {
		return false
	}

	node.Value = value

	return true
}

func (dll *DoublyLinkedList) Insert(index, value int) bool {
	if index == 0 {
		dll.Unshift(value)
		return true
	}

	if index == dll.Length {
		dll.Push(value)
		return true
	}

	node := dll.Get(index)

	if node == nil {
		return false
	}

	node.Previous.Next = &Node{value, node, node.Previous}
	dll.Length++

	return true
}

func (dll *DoublyLinkedList) Remove(index int) *Node {
	if index < 0 || index >= dll.Length {
		return nil
	}

	if index == 0 {
		return dll.Shift()
	}

	if index == dll.Length-1 {
		return dll.Pop()
	}

	node := dll.Get(index)
	node.Previous.Next = node.Next
	node.Previous = nil
	node.Next = nil
	dll.Length--

	return node
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

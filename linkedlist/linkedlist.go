// Package linkedlist provides a basic LinkedList type and methods to handle it, like Push(), Pop(), etc.
package linkedlist

import "fmt"

// Node is each node in the LinkedList
type Node struct {
	Value int
	Next  *Node
}

func (node Node) String() string {
	return fmt.Sprintf("Value: %v -> Next: %v", node.Value, node.Next)
}

// LinkedList has two pointers:
//
// Head points to the first Node.
//
// Tail points to the last Node.
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (ll LinkedList) String() string {
	return fmt.Sprintf("Head: %v / Tail: %v / Length: %v", ll.Head, ll.Tail, ll.Length)
}

// New creates a new LinkedList and sets its first node value.
func New(value int) *LinkedList {
	n := &Node{value, nil}
	ll := &LinkedList{n, n, 1}
	return ll
}

// Push adds a new node to the end of the Linked List.
func (ll *LinkedList) Push(value int) *LinkedList {
	n := &Node{value, nil}

	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		ll.Tail.Next = n
		ll.Tail = n
	}

	ll.Tail.Next = nil
	ll.Length++

	return ll
}

// Pop removes the tail node and returns it.
func (ll *LinkedList) Pop() *Node {
	if ll.Length == 0 {
		return nil
	}

	if ll.Length == 1 {
		t := ll.Head
		ll.Head = nil
		ll.Tail = nil
		ll.Length = 0
		return t
	}

	prev := ll.Head
	curr := ll.Head.Next
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	ll.Tail = prev
	ll.Tail.Next = nil
	curr.Next = nil
	ll.Length--

	return curr
}

// Unshift adds a new item to the beginning of the Linked List.
func (ll *LinkedList) Unshift(value int) *LinkedList {
	temp := ll.Head
	ll.Head = &Node{value, temp}

	if ll.Length == 0 {
		ll.Tail = ll.Head
	}

	ll.Length++

	return ll
}

// Shift removes an items from the beginning of the Linked List.
func (ll *LinkedList) Shift() *Node {
	if ll.Length == 0 {
		return nil
	}

	oldHead := ll.Head
	ll.Head = oldHead.Next
	oldHead.Next = nil
	ll.Length--

	if ll.Length == 0 {
		ll.Tail = nil
	}

	return oldHead
}

// Get returns the Node that is on the index position of the list.
func (ll LinkedList) Get(index int) *Node {
	if index < 0 || index > ll.Length {
		return nil
	}

	if index == 0 {
		return ll.Head
	}

	if index == ll.Length-1 {
		return ll.Tail
	}

	curr := ll.Head.Next
	for i := 1; i < index; i++ {
		curr = curr.Next
	}

	return curr
}

// Set changes the value of a Node on the list and returns true if successful.
func (ll *LinkedList) Set(index int, value int) bool {
	node := ll.Get(index)

	if node == nil {
		return false
	}

	node.Value = value

	return true
}

// Insert adds an item to the specified index and move the other nodes one step to the right
func (ll *LinkedList) Insert(index int, value int) bool {
	if index < 0 || index > ll.Length {
		return false
	}

	if index == 0 {
		ll.Unshift(value)
		return true
	}

	if ll.Length == index {
		ll.Push(value)
		return true
	}

	prevNode := ll.Get(index - 1)
	prevNode.Next = &Node{value, prevNode.Next}
	ll.Length++

	return true
}

// Removes an item from the specified index
func (ll *LinkedList) Remove(index int) *Node {
	if index < 0 || index > ll.Length {
		return nil
	}

	if index == 0 {
		return ll.Shift()
	}

	if index == ll.Length-1 {
		return ll.Pop()
	}

	prevNode := ll.Get(index - 1)
	temp := prevNode.Next
	prevNode.Next = temp.Next
	temp.Next = nil
	ll.Length--

	return temp
}

func (ll *LinkedList) Reverse() {
	if ll.Length <= 1 {
		return
	}

	if ll.Length == 2 {
		ll.Head, ll.Tail = ll.Tail, ll.Head
		ll.Tail.Next = nil
		ll.Head.Next = ll.Tail

		return
	}

	second := ll.Head.Next
	ll.Head, ll.Tail = ll.Tail, ll.Head
	ll.Tail.Next = nil

	a, b, c := second, second.Next, second.Next.Next
	for {
		b.Next = a

		if c == nil {
			break
		}

		a, b, c = b, c, c.Next
	}

	second.Next = ll.Tail
}

// Print shows the details of the Linked List on the screen.
func (ll LinkedList) Print() {
	if ll.Length > 0 {
		currentNode := ll.Head

		for currentNode != nil {
			fmt.Print(currentNode.Value, "->")
			currentNode = currentNode.Next
		}

		fmt.Print("nil")

		fmt.Println()
		fmt.Println("Head:", ll.Head.Value)
		fmt.Println("Tail:", ll.Tail.Value)
	}

	fmt.Println("Length:", ll.Length)
}

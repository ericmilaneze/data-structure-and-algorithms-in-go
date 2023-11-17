package linkedlist

import "testing"

func TestLinkedList_FindKthNodeFromEnd(t *testing.T) {
	ll := New(1)
	ll.Push(2)
	ll.Push(3)
	ll.Push(4)
	ll.Push(5)
	ll.Push(6)

	node := ll.FindKthNodeFromEnd(4)

	if node.Value != 3 {
		t.Error("wrong value\r\nwant: 3\r\ngot:", node.Value)
	}
}

func TestLinkedList_FindKthNodeFromEnd_odd_length(t *testing.T) {
	ll := New(1)
	ll.Push(2)
	ll.Push(3)
	ll.Push(4)
	ll.Push(5)

	node := ll.FindKthNodeFromEnd(2)

	if node.Value != 4 {
		t.Error("wrong value\r\nwant: 4\r\ngot:", node.Value)
	}
}

func TestLinkedList_FindKthNodeFromEnd_k_greater_than_length(t *testing.T) {
	ll := New(1)
	ll.Push(2)
	ll.Push(3)
	ll.Push(4)
	ll.Push(5)

	node := ll.FindKthNodeFromEnd(6)

	if node != nil {
		t.Error("node should be nil")
	}
}

package linkedlist

import (
	"testing"
)

func BenchmarkTestLinkedList_FindMiddleNode(b *testing.B) {
	ll := FromValues(1, 2, 3, 4, 5, 6)

	for i := 0; i < b.N; i++ {
		ll.FindMiddleNode()
	}
}

func TestLinkedList_FindMiddleNode_empty(t *testing.T) {
	ll := new(LinkedList)

	node := ll.FindMiddleNode()

	if node != nil {
		t.Error("ndoe should return nil when the list is empty")
	}
}

func TestLinkedList_FindMiddleNode_one_item(t *testing.T) {
	ll := New(1)

	node := ll.FindMiddleNode()

	if node.Value != 1 {
		t.Error("wrong return\r\nwant: 1, got:", node.Value)
	}
}

func TestLinkedList_FindMiddleNode_odd_length(t *testing.T) {
	ll := FromValues(1, 2, 3, 4, 5)

	node := ll.FindMiddleNode()

	if node.Value != 3 {
		t.Error("wrong return\r\nwant: 3, got:", node.Value)
	}
}

func TestLinkedList_FindMiddleNode_even_length(t *testing.T) {
	ll := FromValues(1, 2, 3, 4, 5, 6)

	node := ll.FindMiddleNode()

	if node.Value != 4 {
		t.Error("wrong return\r\nwant: 4, got:", node.Value)
	}
}

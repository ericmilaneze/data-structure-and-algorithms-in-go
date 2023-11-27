package linkedlist

import "testing"

func TestLinkedList_FindKthNodeFromEnd(t *testing.T) {
	ll := FromValues(1, 2, 3, 4, 5, 6)

	node := ll.FindKthNodeFromEnd(4)

	if node.Value != 3 {
		t.Error("wrong value\nwant: 3\ngot:", node.Value)
	}
}

func TestLinkedList_FindKthNodeFromEnd_odd_length(t *testing.T) {
	ll := FromValues(1, 2, 3, 4, 5)

	node := ll.FindKthNodeFromEnd(2)

	if node.Value != 4 {
		t.Error("wrong value\nwant: 4\ngot:", node.Value)
	}
}

func TestLinkedList_FindKthNodeFromEnd_k_greater_than_length(t *testing.T) {
	ll := FromValues(1, 2, 3, 4, 5)

	node := ll.FindKthNodeFromEnd(6)

	if node != nil {
		t.Error("node should be nil")
	}
}

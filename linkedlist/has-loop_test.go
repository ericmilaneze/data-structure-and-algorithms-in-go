package linkedlist

import "testing"

func TestLinkedList_HasLoop(t *testing.T) {
	ll := New(11)
	ll.Push(3)
	ll.Push(7)
	ll.Tail.Next = ll.Head

	hasLoop := ll.HasLoop()

	if !hasLoop {
		t.Error("it should be true")
	}
}

func TestLinkedList_HasLoop_tail_pointing_to_node_outside_list(t *testing.T) {
	ll := New(11)
	ll.Push(3)
	ll.Push(7)
	ll.Tail.Next = ll.Head

	ll.Tail = &Node{
		Value: 4,
		Next:  nil,
	}

	hasLoop := ll.HasLoop()

	if !hasLoop {
		t.Error("it should be true")
	}
}

func TestLinkedList_HasLoop_not_a_loop(t *testing.T) {
	ll := New(11)
	ll.Push(3)
	ll.Push(7)

	hasLoop := ll.HasLoop()

	if hasLoop {
		t.Error("it should be false")
	}
}

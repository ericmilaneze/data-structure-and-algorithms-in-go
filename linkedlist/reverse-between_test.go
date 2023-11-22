package linkedlist

import "testing"

func TestLinkedList_ReverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		m     int
		n     int
		want  string
		tail  *Node
	}{
		{"empty list", []int{}, 0, 0, "nil", nil},
		{"single-element list", []int{1}, 0, 0, "1->nil", &Node{1, nil}},
		{"reverse includes head", []int{1, 2, 3}, 0, 1, "2->1->3->nil", &Node{3, nil}},
		{"m less than 0", []int{1, 2, 3}, -10, 1, "2->1->3->nil", &Node{3, nil}},
		{"n less than m", []int{1, 2, 3}, 10, 1, "1->2->3->nil", &Node{3, nil}},
		{"reverse includes tail", []int{1, 2, 3}, 1, 2, "1->3->2->nil", &Node{2, nil}},
		{"reverse includes tail (n bigger than length boundary)", []int{1, 2, 3}, 1, 10, "1->3->2->nil", &Node{2, nil}},
		{"reverse between (1)", []int{1, 2, 3, 4, 5}, 1, 3, "1->4->3->2->5->nil", &Node{5, nil}},
		{"reverse between (2)", []int{1, 2, 3, 4, 5, 6}, 3, 5, "1->2->3->6->5->4->nil", &Node{4, nil}},
	}

	for _, test := range tests {
		ll := FromValues(test.items...)

		ll.ReverseBetween(test.m, test.n)

		r := ll.SprintValues()

		if r != test.want {
			t.Errorf("wrong result for \"%s\" (m: %v, n: %v)\r\nwant: %s\r\ngot: %s", test.name, test.m, test.n, test.want, r)
		}

		if ll.Tail == nil && test.tail != nil {
			t.Fatalf("tail shouldn't be nil for \"%s\"\r\nwant: %v\r\ngot: nil", test.name, test.tail)
		}

		if ll.Tail != nil && *ll.Tail != *test.tail {
			t.Errorf("wrong resulting tail for \"%s\"\r\nwant: %v\r\ngot: %v", test.name, *test.tail, *ll.Tail)
		}
	}
}

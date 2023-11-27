package linkedlist

import "testing"

func TestLinkedList_RemoveDuplicates(t *testing.T) {
	tests := []struct {
		name   string
		items  []int
		want   string
		length int
	}{
		{"empty", []int{}, "nil", 0},
		{"one item", []int{1}, "1->nil", 1},
		{"unique elements", []int{1, 2, 3, 4, 5, 0, 11, -1}, "1->2->3->4->5->0->11->-1->nil", 8},
		{"single duplicate", []int{1, 2, 3, 2, 1, 4}, "1->2->3->4->nil", 4},
		{"multiple sets", []int{1, 2, 3, 3, 2, 1, 4}, "1->2->3->4->nil", 4},
		{"duplicates in tail", []int{1, 2, 3, 3, 2, 1, 4, 3, 4}, "1->2->3->4->nil", 4},
		{"all-duplicate", []int{3, 3, 3}, "3->nil", 1},
	}

	for _, test := range tests {
		ll := FromValues(test.items...)

		ll.RemoveDuplicates()

		r := ll.SprintValues()

		if r != test.want {
			t.Errorf("wrong result on test \"%s\"\nwant: %s\ngot: %s", test.name, test.want, r)
		}

		if ll.Length != test.length {
			t.Errorf("wrong length on test \"%s\"\nwant: %v\ngot: %v", test.name, test.length, ll.Length)
		}
	}
}

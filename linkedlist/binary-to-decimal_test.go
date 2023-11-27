package linkedlist

import "testing"

func TestLinkedList_BinaryToDecimal(t *testing.T) {
	tests := []struct {
		items []int
		want  int
	}{
		{items: []int{}, want: 0},
		{items: []int{1}, want: 1},
		{items: []int{0}, want: 0},
		{items: []int{0, 1, 0}, want: 2},
		{items: []int{1, 1, 0}, want: 6},
		{items: []int{1, 1, 1}, want: 7},
		{items: []int{1, 0, 0, 0}, want: 8},
	}

	for _, test := range tests {
		ll := FromValues(test.items...)

		r := ll.BinaryToDecimal()

		if r != test.want {
			t.Errorf("wrong return for \"%s\"\nwant: %v\ngot: %v", ll.SprintValues(), test.want, r)
		}
	}
}

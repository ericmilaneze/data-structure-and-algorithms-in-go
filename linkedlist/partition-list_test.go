package linkedlist

import "testing"

func TestLinkedList_PartitionList_empty(t *testing.T) {
	ll := new(LinkedList)

	ll.PartitionList(0)

	if ll.Length != 0 {
		t.Error("wrong length\nwant: 0\ngot:", ll.Length)
	}
}

func TestLinkedList_PartitionList(t *testing.T) {
	tests := []struct {
		name   string
		x      int
		values []int
		want   string
		tail   int
	}{
		{name: "single element", x: 0, values: []int{1}, want: "1->nil", tail: 1},
		{name: "all values less than x", x: 15, values: []int{6, 7, 8, 9, 10}, want: "6->7->8->9->10->nil", tail: 10},
		{name: "all values greater than x", x: 15, values: []int{26, 27, 28, 29, 30}, want: "26->27->28->29->30->nil", tail: 30},
		{name: "mixed values 1", x: 15, values: []int{6, 26, 27, 7, 8, 28, 9, 30, 10}, want: "6->7->8->9->10->26->27->28->30->nil", tail: 30},
		{name: "mixed values 2", x: 15, values: []int{26, 6, 7, 27, 8, 28, 9, 29, 30, 10, 31}, want: "6->7->8->9->10->26->27->28->29->30->31->nil", tail: 31},
	}

	for _, test := range tests {
		ll := FromValues(test.values...)

		ll.PartitionList(test.x)

		ret := ll.SprintValues()

		if ret != test.want {
			t.Errorf("wrong return on test \"%s\"\nwant: %v\ngot: %v", test.name, test.want, ret)
		}

		if ll.Tail.Value != test.tail {
			t.Errorf("wrong tail on test \"%s\"\nwant: %v\ngot: %v", test.name, test.tail, ll.Tail.Value)
		}
	}
}

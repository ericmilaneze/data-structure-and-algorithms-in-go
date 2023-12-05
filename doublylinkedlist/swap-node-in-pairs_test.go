package doublylinkedlist

import "testing"

func TestDoublyLinkedList_SwapPairs(t *testing.T) {
	type want struct {
		result    string
		length    int
		headValue int
		tailValue int
	}

	tests := []struct {
		name  string
		items []int
		want  want
	}{
		{"empty list", []int{}, want{"nil", 0, 0, 0}},
		{"one element", []int{1}, want{"nil<-1->nil", 1, 1, 1}},
		{"two elements", []int{1, 2}, want{"nil<-2<-1->nil", 2, 2, 1}},
		{"three elements", []int{1, 2, 3}, want{"nil<-2<-1<-3->nil", 2, 2, 3}},
		{"four elements", []int{1, 2, 3, 4}, want{"nil<-2<-1<-4->3->nil", 2, 2, 3}},
		{"long list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want{"nil<-2<->1<->4<->3<->6<->5<->8<->7<->9->nil", 9, 2, 9}},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		dll.SwapPairs()

		if dll.SprintValues() != test.want.result {
			t.Errorf("test \"%s\" failed - the resulting list is wrong\nwant: %s\ngot: %s", test.name, test.want.result, dll.SprintValues())
		}

		if dll.Length != test.want.length {
			t.Fatalf("test \"%s\" failed - the resulting length is wrong\nwant: %v\ngot: %v", test.name, test.want.length, dll.Length)
		}

		if test.want.length == 0 {
			if dll.Head != nil {
				t.Errorf("test \"%s\" failed - an empty list should not have a Head", test.name)
			}

			if dll.Tail != nil {
				t.Errorf("test \"%s\" failed - an empty list should not have a Tail", test.name)
			}
		} else {
			if dll.Head.Value != test.want.headValue {
				t.Errorf("test \"%s\" failed - Head value is wrong\nwant: %v\ngot: %v", test.name, test.want.headValue, dll.Head.Value)
			}

			if dll.Tail.Value != test.want.tailValue {
				t.Errorf("test \"%s\" failed - Tail value is wrong\nwant: %v\ngot: %v", test.name, test.want.tailValue, dll.Tail.Value)
			}
		}
	}
}

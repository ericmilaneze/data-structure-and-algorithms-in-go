package doublylinkedlist

import "testing"

func TestDoublyLinkedList_Reverse(t *testing.T) {
	tests := []struct {
		name      string
		items     []int
		result    string
		length    int
		headValue int
		tailValue int
	}{
		{"empty list", []int{}, "nil", 0, 0, 0},
		{"single element", []int{1}, "nil<-1->nil", 1, 1, 1},
		{"two elements", []int{1, 2}, "nil<-2<->1->nil", 2, 2, 1},
		{"three elements", []int{1, 2, 3}, "nil<-3<->2<->1->nil", 3, 3, 1},
		{"even-length list", []int{1, 2, 3, 4, 5, 6}, "nil<-6<->5<->4<->3<->2<->1->nil", 6, 6, 1},
		{"odd-length list", []int{1, 2, 3, 4, 5, 6, 7}, "nil<-7<->6<->5<->4<->3<->2<->1->nil", 7, 7, 1},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		dll.Reverse()

		if dll.SprintValues() != test.result {
			t.Errorf("test \"%s\" failed - result is different than expected\nwant: %v\ngot: %v", test.name, test.result, dll.SprintValues())
		}

		if dll.Length != test.length {
			t.Fatalf("test \"%s\" failed - length is different than expected\nwant: %v\ngot: %v", test.name, test.length, dll.Length)
		}

		if dll.Length > 0 {
			if dll.Head == nil {
				t.Fatalf("test \"%s\" failed - Head shouldn't be nil", test.name)
			}

			if dll.Head.Value != test.headValue {
				t.Errorf("test \"%s\" failed - Head is different than expected\nwant: %v\ngot: %v", test.name, test.headValue, dll.Head.Value)
			}

			if dll.Tail == nil {
				t.Fatalf("test \"%s\" failed - Tail shouldn't be nil", test.name)
			}

			if dll.Tail.Value != test.tailValue {
				t.Errorf("test \"%s\" failed - Tail is different than expected\nwant: %v\ngot: %v", test.name, test.tailValue, dll.Tail.Value)
			}
		}
	}
}

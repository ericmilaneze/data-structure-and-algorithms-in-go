package doublylinkedlist

import (
	"fmt"
	"testing"
)

func ExampleNode() {
	node := &Node{
		Value:    1,
		Next:     &Node{Value: 2},
		Previous: &Node{Value: 0},
	}

	fmt.Print(node)

	// Output:
	// Value: 1
	// Next: 2
	// Previous: 0
}

func ExampleDoublyLinkedList() {
	dll := FromValues(10, 11, 12, 13, 14)
	fmt.Print(dll)

	// Output:
	// Head: 10
	// Tail: 14
	// Length: 5
	// Items: nil<-10<->11<->12<->13<->14->nil
	// Items (reverse): nil<-14<->13<->12<->11<->10->nil
}

func ExampleDoublyLinkedList_one_item() {
	dll := New(10)
	fmt.Print(dll)

	// Output:
	// Head: 10
	// Tail: 10
	// Length: 1
	// Items: nil<-10->nil
	// Items (reverse): nil<-10->nil
}

func ExampleDoublyLinkedList_empty() {
	dll := new(DoublyLinkedList)
	fmt.Print(dll)

	// Output:
	// Length: 0
	// Items: nil
	// Items (reverse): nil
}

func TestDoublyLinkedList_Pop(t *testing.T) {
	type want struct {
		value     int
		length    int
		headValue int
		tailValue int
	}

	tests := []struct {
		name  string
		items []int
		want  want
		empty bool
	}{
		{name: "empty list", items: []int{}, want: want{0, 0, 0, 0}, empty: true},
		{name: "single node", items: []int{1}, want: want{1, 0, 0, 0}, empty: true},
		{name: "multiple nodes", items: []int{1, 2, 3, 4, 5}, want: want{5, 4, 1, 4}},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		r := dll.Pop()

		if dll.Length != test.want.length {
			t.Errorf("test \"%s\" - length is wrong\nwant: %v\ngot: %v", test.name, test.want.length, dll.Length)
		}

		if test.empty {
			if dll.Head != nil {
				t.Errorf("test \"%s\" was expecting a nil Head", test.name)
			}

			if dll.Tail != nil {
				t.Errorf("test \"%s\" was expecting a nil Tail", test.name)
			}

			t.Logf("test \"%s\" - empty list hence skipping the other tests", test.name)
			continue
		}

		if r == nil {
			t.Fatalf("test \"%s\" wasn't expecting nil as return", test.name)
		}

		if r.Value != test.want.value {
			t.Errorf("test \"%s\" - wrong value returned\nwant: %v\ngot: %v", test.name, test.want.value, r.Value)
		}

		if dll.Head.Value != test.want.headValue {
			t.Errorf("test \"%s\" - wrong head value\nwant: %v\ngot: %v", test.name, test.want.headValue, dll.Head.Value)
		}

		if dll.Tail.Value != test.want.tailValue {
			t.Errorf("test \"%s\" - wrong tail value\nwant: %v\ngot: %v", test.name, test.want.tailValue, dll.Tail.Value)
		}
	}
}

func TestLinkedList_Unshift(t *testing.T) {
	type want struct {
		length    int
		headValue int
		tailValue int
		result    string
	}

	tests := []struct {
		name  string
		value int
		items []int
		want  want
	}{
		{"empty list", 10, []int{}, want{1, 10, 10, "nil<-10->nil"}},
		{"single node", 10, []int{1}, want{2, 10, 1, "nil<-10<->1->nil"}},
		{"multiple nodes", 10, []int{1, 2, 3, 4, 5}, want{6, 10, 5, "nil<-10<->1<->2<->3<->4<->5->nil"}},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		dll.Unshift(test.value)

		if dll.Length != test.want.length {
			t.Errorf("test \"%s\" - length is wrong\nwant: %v\ngot: %v", test.name, test.want.length, dll.Length)
		}

		if dll.Head.Value != test.want.headValue {
			t.Errorf("test \"%s\" - wrong head value\nwant: %v\ngot: %v", test.name, test.want.headValue, dll.Head.Value)
		}

		if dll.Tail.Value != test.want.tailValue {
			t.Errorf("test \"%s\" - wrong tail value\nwant: %v\ngot: %v", test.name, test.want.tailValue, dll.Tail.Value)
		}

		if resultingList := dll.SprintValues(); resultingList != test.want.result {
			t.Errorf("test \"%s\" - resulting list is wrong\nwant: %v\ngot: %v", test.name, test.want.result, resultingList)
		}
	}
}

func TestDoublyLinkedList_Shift(t *testing.T) {
	type want struct {
		value     int
		length    int
		headValue int
		tailValue int
	}

	tests := []struct {
		name  string
		items []int
		want  want
		empty bool
	}{
		{"empty list", []int{}, want{0, 0, 0, 0}, true},
		{"single node", []int{1}, want{1, 0, 0, 0}, true},
		{"multiple nodes", []int{1, 2, 3, 4, 5}, want{1, 4, 2, 5}, false},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		r := dll.Shift()

		if dll.Length != test.want.length {
			t.Errorf("test \"%s\" - length is wrong\nwant: %v\ngot: %v", test.name, test.want.length, dll.Length)
		}

		if test.empty {
			if dll.Head != nil {
				t.Errorf("test \"%s\" was expecting a nil Head", test.name)
			}

			if dll.Tail != nil {
				t.Errorf("test \"%s\" was expecting a nil Tail", test.name)
			}

			t.Logf("test \"%s\" - empty list hence skipping the other tests", test.name)
			continue
		}

		if r == nil {
			t.Fatalf("test \"%s\" wasn't expecting nil as return", test.name)
		}

		if r.Value != test.want.value {
			t.Errorf("test \"%s\" - wrong value returned\nwant: %v\ngot: %v", test.name, test.want.value, r.Value)
		}

		if dll.Head.Value != test.want.headValue {
			t.Errorf("test \"%s\" - wrong head value\nwant: %v\ngot: %v", test.name, test.want.headValue, dll.Head.Value)
		}

		if dll.Tail.Value != test.want.tailValue {
			t.Errorf("test \"%s\" - wrong tail value\nwant: %v\ngot: %v", test.name, test.want.tailValue, dll.Tail.Value)
		}
	}
}

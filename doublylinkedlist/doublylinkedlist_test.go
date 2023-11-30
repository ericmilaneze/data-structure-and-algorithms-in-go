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

func TestDoublyLinkedList_Get(t *testing.T) {
	tests := []struct {
		name       string
		items      []int
		index      int
		returnsNil bool
		want       int
	}{
		{"negative index", []int{1, 2, 3, 4, 5}, -1, true, 0},
		{"index greater than length", []int{1, 2, 3, 4, 5}, 5, true, 0},
		{"return head for index 0", []int{1, 2, 3, 4, 5}, 0, false, 1},
		{"return tail for index length -1", []int{1, 2, 3, 4, 5}, 4, false, 5},
		{"middle index", []int{1, 2, 3, 4, 5}, 2, false, 3},
		{"should start from head for index < length/2", []int{1, 2, 3, 4, 5}, 1, false, 2},
		{"should start from tail for index > length/2", []int{1, 2, 3, 4, 5}, 3, false, 4},
		{"single element", []int{1}, 0, false, 1},
		{"empty list", []int{}, 0, true, 0},
		{"index = length/2", []int{1, 2, 3, 4, 5, 6}, 3, false, 4},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		r := dll.Get(test.index)

		if test.returnsNil {
			if r != nil {
				t.Errorf("test \"%s\" - was expecting nil", test.name)
			}

			continue
		}

		if r == nil {
			t.Fatalf("test \"%s\" - wasn't expecting nil\nwant: %v\ngot: nil", test.name, test.want)
		}

		if r.Value != test.want {
			t.Errorf("test \"%s\" - wrong value returned\nwant: %v\ngot: %v", test.name, test.want, r.Value)
		}
	}
}

func TestDoublyLinkedList_Set(t *testing.T) {
	type want struct {
		returns bool
		result  string
	}

	tests := []struct {
		name  string
		items []int
		index int
		value int
		want  want
	}{
		{"negative index", []int{1, 2, 3, 4, 5}, -1, 0, want{false, "nil<-1<->2<->3<->4<->5->nil"}},
		{"index greater than length", []int{1, 2, 3, 4, 5}, 5, 0, want{false, "nil<-1<->2<->3<->4<->5->nil"}},
		{"set head when index = 0", []int{1, 2, 3, 4, 5}, 0, 10, want{true, "nil<-10<->2<->3<->4<->5->nil"}},
		{"set tail for index = length -1", []int{1, 2, 3, 4, 5}, 4, 50, want{true, "nil<-1<->2<->3<->4<->50->nil"}},
		{"middle index", []int{1, 2, 3, 4, 5}, 2, 30, want{true, "nil<-1<->2<->30<->4<->5->nil"}},
		{"index < length/2", []int{1, 2, 3, 4, 5}, 1, 20, want{true, "nil<-1<->20<->3<->4<->5->nil"}},
		{"index > length/2", []int{1, 2, 3, 4, 5}, 3, 40, want{true, "nil<-1<->2<->3<->40<->5->nil"}},
		{"single element", []int{1}, 0, 10, want{true, "nil<-10->nil"}},
		{"empty list", []int{}, 0, 0, want{false, "nil"}},
		{"index = length/2", []int{1, 2, 3, 4, 5, 6}, 3, 40, want{true, "nil<-1<->2<->3<->40<->5<->6->nil"}},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		r := dll.Set(test.index, test.value)

		if r != test.want.returns {
			t.Errorf("test \"%s\" - wrong return\nwant: %v\ngot: %v", test.name, test.want.returns, r)
		}

		if dll.SprintValues() != test.want.result {
			t.Errorf("test \"%s\" - wrong result\nwant: %v\ngot: %v", test.name, test.want.result, dll.SprintValues())
		}
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	type want struct {
		returns   bool
		result    string
		headValue int
		tailValue int
	}

	tests := []struct {
		name  string
		items []int
		index int
		value int
		want  want
	}{
		{"negative index", []int{1, 2, 3, 4, 5}, -1, 0, want{false, "nil<-1<->2<->3<->4<->5->nil", 1, 5}},
		{"index greater than length", []int{1, 2, 3, 4, 5}, 10, 0, want{false, "nil<-1<->2<->3<->4<->5->nil", 1, 5}},
		{"index = length + 1", []int{1, 2, 3, 4, 5}, 5, 6, want{true, "nil<-1<->2<->3<->4<->5<->6->nil", 1, 6}},
		{"index = 0", []int{1, 2, 3, 4, 5}, 0, 10, want{true, "nil<-10<->1<->2<->3<->4<->5->nil", 10, 5}},
		{"index = length -1", []int{1, 2, 3, 4, 5}, 4, 50, want{true, "nil<-1<->2<->3<->4<->50<->5->nil", 1, 5}},
		{"middle index", []int{1, 2, 3, 4, 5}, 2, 30, want{true, "nil<-1<->2<->30<->3<->4<->5->nil", 1, 5}},
		{"index < length/2", []int{1, 2, 3, 4, 5}, 1, 20, want{true, "nil<-1<->20<->2<->3<->4<->5->nil", 1, 5}},
		{"index > length/2", []int{1, 2, 3, 4, 5}, 3, 40, want{true, "nil<-1<->2<->3<->40<->4<->5->nil", 1, 5}},
		{"single element", []int{1}, 0, 10, want{true, "nil<-10<->1->nil", 10, 1}},
		{"empty list", []int{}, 0, 1, want{true, "nil<-1->nil", 1, 1}},
		{"index = length/2", []int{1, 2, 3, 4, 5, 6}, 3, 40, want{true, "nil<-1<->2<->3<->40<->4<->5<->6->nil", 1, 6}},
	}

	for _, test := range tests {
		dll := new(DoublyLinkedList)

		for i, v := range test.items {
			dll.Insert(i, v)
		}

		r := dll.Insert(test.index, test.value)

		if r != test.want.returns {
			t.Errorf("test \"%s\" - wrong return\nwant: %v\ngot: %v", test.name, test.want.returns, r)
		}

		if dll.SprintValues() != test.want.result {
			t.Errorf("test \"%s\" - wrong result\nwant: %v\ngot: %v", test.name, test.want.result, dll.SprintValues())
		}

		if dll.Head == nil {
			t.Fatal("head shouldn't be nil after inserting a new value")
		}

		if dll.Head.Value != test.want.headValue {
			t.Errorf("test \"%s\" - wrong head value\nwant: %v\ngot: %v", test.name, test.want.headValue, dll.Head.Value)
		}

		if dll.Tail == nil {
			t.Fatal("head shouldn't be nil after inserting a new value")
		}

		if dll.Tail.Value != test.want.tailValue {
			t.Errorf("test \"%s\" - wrong tail value\nwant: %v\ngot: %v", test.name, test.want.tailValue, dll.Tail.Value)
		}
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	type want struct {
		shouldReturn bool
		returns      int
		result       string
		headValue    int
		tailValue    int
	}

	tests := []struct {
		name  string
		items []int
		index int
		want  want
	}{
		{"negative index", []int{1, 2, 3, 4, 5}, -1, want{false, 0, "nil<-1<->2<->3<->4<->5->nil", 1, 5}},
		{"index greater than length", []int{1, 2, 3, 4, 5}, 10, want{false, 0, "nil<-1<->2<->3<->4<->5->nil", 1, 5}},
		{"index = length + 1", []int{1, 2, 3, 4, 5}, 5, want{false, 0, "nil<-1<->2<->3<->4<->5->nil", 1, 5}},
		{"index = 0", []int{1, 2, 3, 4, 5}, 0, want{true, 1, "nil<-2<->3<->4<->5->nil", 2, 5}},
		{"index = length -1", []int{1, 2, 3, 4, 5}, 4, want{true, 5, "nil<-1<->2<->3<->4->nil", 1, 4}},
		{"middle index", []int{1, 2, 3, 4, 5}, 2, want{true, 3, "nil<-1<->2<->4<->5->nil", 1, 5}},
		{"index < length/2", []int{1, 2, 3, 4, 5}, 1, want{true, 2, "nil<-1<->3<->4<->5->nil", 1, 5}},
		{"index > length/2", []int{1, 2, 3, 4, 5}, 3, want{true, 4, "nil<-1<->2<->3<->5->nil", 1, 5}},
		{"single element", []int{1}, 0, want{true, 1, "nil", 0, 0}},
		{"empty list", []int{}, 0, want{false, 0, "nil", 0, 0}},
		{"index = length/2", []int{1, 2, 3, 4, 5, 6}, 3, want{true, 4, "nil<-1<->2<->3<->5<->6->nil", 1, 6}},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		r := dll.Remove(test.index)

		if test.want.shouldReturn {
			if r == nil {
				t.Fatalf("test \"%s\" - should not return nil\nwant: %v\ngot: nil", test.name, test.want.returns)
			}

			if r.Value != test.want.returns {
				t.Errorf("test \"%s\" - wrong return\nwant: %v\ngot: %v", test.name, test.want.returns, r.Value)
			}
		} else {
			if r != nil {
				t.Errorf("test \"%s\" - should return nil\nwant: nil\ngot: %v", test.name, r.Value)
			}
		}

		if dll.SprintValues() != test.want.result {
			t.Errorf("test \"%s\" - wrong result\nwant: %v\ngot: %v", test.name, test.want.result, dll.SprintValues())
		}

		if dll.Head == nil {
			t.Fatal("head shouldn't be nil after inserting a new value")
		}

		if dll.Head.Value != test.want.headValue {
			t.Errorf("test \"%s\" - wrong head value\nwant: %v\ngot: %v", test.name, test.want.headValue, dll.Head.Value)
		}

		if dll.Tail == nil {
			t.Fatal("head shouldn't be nil after inserting a new value")
		}

		if dll.Tail.Value != test.want.tailValue {
			t.Errorf("test \"%s\" - wrong tail value\nwant: %v\ngot: %v", test.name, test.want.tailValue, dll.Tail.Value)
		}
	}
}

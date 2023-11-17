package linkedlist

import (
	"fmt"
	"testing"
)

func Example() {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)

	fmt.Println(ll)
	ll.Print()

	// Output:
	// Head: Value: 10 -> Next: Value: 11 -> Next: Value: 12 -> Next: Value: 13 -> Next: Value: 14 -> Next: <nil> / Tail: Value: 14 -> Next: <nil> / Length: 5
	// 10->11->12->13->14->nil
	// Head: 10
	// Tail: 14
	// Length: 5
}

func ExampleNew() {
	ll := New(10)

	fmt.Println(ll.Head.Value)
	fmt.Println(ll.Tail.Value)
	fmt.Println(ll.Length)

	// Output:
	// 10
	// 10
	// 1
}

func TestNew(t *testing.T) {
	v := 10
	length := 1

	ll := New(v)

	if ll.Length != length {
		t.Fatal("length\r\nwant:", length, "\r\ngot:", ll.Length)
	}

	if ll.Head == nil {
		t.Fatal("head wasn't defined")
	}

	if ll.Head.Value != v {
		t.Error("head value\r\nwant:", v, "\r\ngot:", ll.Head.Value)
	}

	if ll.Tail == nil {
		t.Fatal("tail wasn't defined")
	}

	if ll.Tail.Value != v {
		t.Error("tail value\r\nwant:", v, "\r\ngot:", ll.Tail.Value)
	}

	if ll.Tail.Next != nil {
		t.Error("tail should point to nil")
	}
}

func ExampleLinkedList_Push() {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)

	fmt.Println(ll.Head.Value)
	fmt.Println(ll.Tail.Value)
	fmt.Println(ll.Length)

	// Output:
	// 10
	// 14
	// 5
}

func TestLinkedList_Push_one_item(t *testing.T) {
	vi := 1
	vp := 2
	length := 2

	ll := New(vi)
	r := ll.Push(vp)

	if r != ll {
		t.Fatal("wrong pointer was returned")
	}

	if ll.Length != length {
		t.Error("length\r\nwant:", length, "\r\ngot:", ll.Length)
	}

	if ll.Head == nil {
		t.Fatal("head wasn't defined")
	}

	if ll.Head.Value != vi {
		t.Error("head value\r\nwant:", vi, "\r\ngot:", ll.Head.Value)
	}

	if ll.Tail == nil {
		t.Fatal("tail wasn't defined")
	}

	if ll.Tail.Value != vp {
		t.Error("tail value\r\nwant:", vp, "\r\ngot:", ll.Tail.Value)
	}

	if ll.Tail.Next != nil {
		t.Error("tail should point to nil")
	}
}

func TestLinkedList_Push_one_item_to_an_empty_linked_list(t *testing.T) {
	vp := 2
	length := 1

	ll := &LinkedList{nil, nil, 0}
	r := ll.Push(vp)

	if r != ll {
		t.Fatal("wrong pointer was returned")
	}

	if ll.Length != length {
		t.Error("length\r\nwant:", length, "\r\ngot:", ll.Length)
	}

	if ll.Head == nil {
		t.Fatal("head wasn't defined")
	}

	if ll.Head.Value != vp {
		t.Error("head value\r\nwant:", vp, "\r\ngot:", ll.Head.Value)
	}

	if ll.Tail == nil {
		t.Fatal("tail wasn't defined")
	}

	if ll.Tail.Value != vp {
		t.Error("tail value\r\nwant:", vp, "\r\ngot:", ll.Tail.Value)
	}

	if ll.Tail.Next != nil {
		t.Error("tail should point to nil")
	}
}

func TestLinkedList_Pop(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)

	n := ll.Pop()

	if n == nil {
		t.Fatal("it should return the last value of the Linked List")
	}

	if n.Value != 14 {
		t.Error("popped value\r\nwant: 14\r\ngot:", n.Value)
	}

	if n.Next != nil {
		t.Error("next item should be nil")
	}

	if ll.Length != 4 {
		t.Error("list length\r\nwant:", 4, "\r\ngot:", ll.Length)
	}

	if ll.Tail == nil {
		t.Fatal("tail should not be nil")
	}

	if ll.Tail.Value != 13 {
		t.Error("tail value\r\nwant: 13\r\ngot:", ll.Tail.Value)
	}
}

func TestLinkedList_Pop_all_items(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)

	ll.Pop()
	ll.Pop()
	ll.Pop()
	ll.Pop()
	ll.Pop()
	n := ll.Pop()

	if n != nil {
		t.Error("it should return nil")
	}

	if ll.Length != 0 {
		t.Error("list should be empty")
	}

	if ll.Tail != nil {
		t.Error("tail should be nil")
	}

	if ll.Head != nil {
		t.Error("head should be nil")
	}
}

func TestLinkedList_Unshift_empty_linked_list(t *testing.T) {
	ll := &LinkedList{}

	r := ll.Unshift(9)

	if r != ll {
		t.Fatal("wrong pointer was returned")
	}

	if ll.Head == nil {
		t.Fatal("head wasn't defined")
	}

	if ll.Head.Value != 9 {
		t.Error("head value\r\nwant:", 9, "\r\ngot:", ll.Head.Value)
	}

	if ll.Length != 1 {
		t.Error("length\r\nwant:", 1, "\r\ngot:", ll.Length)
	}

	if ll.Tail == nil {
		t.Fatal("tail wasn't defined")
	}

	if ll.Tail.Value != 9 {
		t.Error("tail value\r\nwant:", 9, "\r\ngot:", ll.Tail.Value)
	}
}

func TestLinkedList_Unshift_one_item(t *testing.T) {
	ll := New(10)

	r := ll.Unshift(9)

	if r != ll {
		t.Fatal("wrong pointer was returned")
	}

	if ll.Head == nil {
		t.Fatal("head wasn't defined")
	}

	if ll.Head.Value != 9 {
		t.Error("head value\r\nwant:", 9, "\r\ngot:", ll.Head.Value)
	}

	if ll.Length != 2 {
		t.Error("length\r\nwant:", 2, "\r\ngot:", ll.Length)
	}

	if ll.Tail == nil {
		t.Fatal("tail wasn't defined")
	}

	if ll.Tail.Value != 10 {
		t.Error("tail value\r\nwant:", 10, "\r\ngot:", ll.Tail.Value)
	}
}

func ExampleLinkedList_Unshift() {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)
	ll.Push(15)

	ll.Unshift(9)

	ll.Print()

	// Output:
	// 9->10->11->12->13->14->15->nil
	// Head: 9
	// Tail: 15
	// Length: 7
}

func TestLinkedList_Unshift_multiple_items(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)
	ll.Push(15)

	r := ll.Unshift(9)

	if r != ll {
		t.Fatal("wrong pointer was returned")
	}

	if ll.Head == nil {
		t.Fatal("head wasn't defined")
	}

	if ll.Head.Value != 9 {
		t.Error("head value\r\nwant:", 9, "\r\ngot:", ll.Head.Value)
	}

	if ll.Length != 7 {
		t.Error("length\r\nwant:", 7, "\r\ngot:", ll.Length)
	}

	if ll.Tail == nil {
		t.Fatal("tail wasn't defined")
	}

	if ll.Tail.Value != 15 {
		t.Error("tail value\r\nwant:", 15, "\r\ngot:", ll.Tail.Value)
	}
}

func TestLinkedList_Shift_empty_linked_list(t *testing.T) {
	ll := &LinkedList{}

	r := ll.Shift()

	if r != nil {
		t.Error("it should return nil")
	}

	if ll.Length != 0 {
		t.Error("length\r\nwant:", 0, "\r\ngot:", ll.Length)
	}

	if ll.Head != nil {
		t.Error("head shouldn't be defined")
	}

	if ll.Tail != nil {
		t.Error("tail shouldn't be defined")
	}
}

func TestLinkedList_Shift_one_item(t *testing.T) {
	ll := New(10)

	r := ll.Shift()

	if r == nil {
		t.Fatal("it shouldn't return nil")
	}

	if r.Value != 10 {
		t.Error("value\r\nwant:", 10, "\r\ngot:", r.Value)
	}

	if ll.Length != 0 {
		t.Error("length\r\nwant:", 0, "\r\ngot:", ll.Length)
	}

	if ll.Head != nil {
		t.Error("head shouldn't be defined")
	}

	if ll.Tail != nil {
		t.Error("tail shouldn't be defined")
	}
}

func ExampleLinkedList_Shift() {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)
	ll.Push(15)

	r := ll.Shift()

	fmt.Println(r.Value)
	ll.Print()

	// Output:
	// 10
	// 11->12->13->14->15->nil
	// Head: 11
	// Tail: 15
	// Length: 5
}

func TestLinkedList_Shift_multiple_items(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)
	ll.Push(15)

	r := ll.Shift()

	if r == nil {
		t.Fatal("it shouldn't return nil")
	}

	if r.Value != 10 {
		t.Error("value\r\nwant:", 10, "\r\ngot:", r.Value)
	}

	if ll.Length != 5 {
		t.Error("length\r\nwant:", 5, "\r\ngot:", ll.Length)
	}

	if ll.Head == nil {
		t.Fatal("head should be defined")
	}

	if ll.Head.Value != 11 {
		t.Error("head value\r\nwant: 11\r\ngot:", ll.Head.Value)
	}

	if ll.Tail == nil {
		t.Fatal("tail should be defined")
	}

	if ll.Tail.Value != 15 {
		t.Error("tail value\r\nwant: 15\r\ngot:", ll.Head.Value)
	}
}

func TestLinkedList_Get(t *testing.T) {
	tests := []struct {
		pos   int
		value int
	}{
		{0, 10},
		{1, 11},
		{2, 12},
	}

	ll := &LinkedList{}

	for _, v := range tests {
		ll.Push(v.value)
	}

	for _, v := range tests {
		r := ll.Get(v.pos)

		if r == nil {
			t.Fatal("wasn't expecting nil")
		}

		if r.Value != v.value {
			t.Error("wrong value\r\npos:", v.pos, "\r\nwant:", v.value, "\r\ngot:", r.Value)
		}
	}
}

func TestLinkedList_Get_empty(t *testing.T) {
	ll := &LinkedList{}

	r := ll.Get(0)

	if r != nil {
		t.Error("should return nil")
	}
}

func TestLinkedList_Get_index_less_than_zero(t *testing.T) {
	ll := &LinkedList{}

	r := ll.Get(-1)

	if r != nil {
		t.Error("should return nil")
	}
}

func TestLinkedList_Get_index_greater_than_length(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)

	r := ll.Get(3)

	if r != nil {
		t.Error("should return nil")
	}
}

func TestLinkedList_Set(t *testing.T) {
	tests := []struct {
		pos      int
		oldValue int
		newValue int
	}{
		{0, 10, 110},
		{1, 11, 111},
		{2, 12, 112},
	}

	ll := &LinkedList{}

	for _, v := range tests {
		ll.Push(v.oldValue)
	}

	for _, v := range tests {
		r := ll.Set(v.pos, v.newValue)

		if r == false {
			t.Error("wasn't expecting false")
		}

		nr := ll.Get(v.pos)

		if nr.Value != v.newValue {
			t.Error("wrong value\r\npos:", v.pos, "\r\nwant:", v.newValue, "\r\ngot:", nr.Value)
		}
	}
}

func TestLinkedList_Set_empty(t *testing.T) {
	ll := &LinkedList{}

	r := ll.Set(0, 10)

	if r != false {
		t.Error("should return false")
	}
}

func TestLinkedList_Set_index_less_than_zero(t *testing.T) {
	ll := &LinkedList{}

	r := ll.Set(-1, 10)

	if r != false {
		t.Error("should return false")
	}
}

func TestLinkedList_Set_index_greater_than_length(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)

	r := ll.Set(3, 10)

	if r != false {
		t.Error("should return false")
	}
}

func ExampleLinkedList_Insert() {
	ll := New(10)
	ll.Push(11)
	ll.Push(13)
	ll.Push(14)
	ll.Push(15)

	r := ll.Insert(2, 12)

	fmt.Println(r)
	ll.Print()

	// Output:
	// true
	// 10->11->12->13->14->15->nil
	// Head: 10
	// Tail: 15
	// Length: 6
}

func TestLinkedList_Insert(t *testing.T) {
	tests := []struct {
		pos      int
		oldValue int
		newValue int
	}{
		{0, 10, 110},
		{1, 11, 111},
		{2, 12, 112},
	}

	ll := &LinkedList{}

	for _, v := range tests {
		r := ll.Insert(v.pos, v.newValue)

		if r == false {
			t.Error("wasn't expecting false")
		}

		nr := ll.Get(v.pos)

		if nr.Value != v.newValue {
			t.Error("wrong value\r\npos:", v.pos, "\r\nwant:", v.newValue, "\r\ngot:", nr.Value)
		}
	}
}

func TestLinkedList_Insert_index_less_than_zero(t *testing.T) {
	ll := &LinkedList{}

	r := ll.Insert(-1, 10)

	if r != false {
		t.Error("should return false")
	}
}

func TestLinkedList_Insert_index_greater_than_length(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)

	r := ll.Insert(4, 10)

	if r != false {
		t.Error("should return false")
	}
}

func ExampleLinkedList_Remove() {
	ll := New(10)
	ll.Push(11)
	ll.Push(121) // will be removed
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)
	ll.Push(15)

	r := ll.Remove(2)

	fmt.Println(r.Value)
	ll.Print()

	// Output:
	// 121
	// 10->11->12->13->14->15->nil
	// Head: 10
	// Tail: 15
	// Length: 6
}

func TestLinkedList_Remove_from_start_to_end(t *testing.T) {
	tests := []int{10, 11, 12, 13, 14, 15}

	ll := &LinkedList{}

	for _, v := range tests {
		ll.Push(v)
	}

	for _, v := range tests {
		r := ll.Remove(0)

		if r == nil {
			t.Error("wasn't expecting nil")
		}

		if r.Value != v {
			t.Error("wrong value\r\nwant:", v, "\r\ngot:", r.Value)
		}
	}
}

func TestLinkedList_Remove_from_end_to_start(t *testing.T) {
	tests := []int{10, 11, 12, 13, 14, 15}

	ll := &LinkedList{}

	for _, v := range tests {
		ll.Push(v)
	}

	for i := len(tests) - 1; i >= 0; i-- {
		r := ll.Remove(ll.Length - 1)

		if r == nil {
			t.Error("wasn't expecting nil")
		}

		if r.Value != tests[i] {
			t.Error("wrong value\r\nwant:", tests[i], "\r\ngot:", r.Value)
		}
	}
}

func TestLinkedList_Remove_index_less_than_zero(t *testing.T) {
	ll := &LinkedList{}

	r := ll.Remove(-1)

	if r != nil {
		t.Error("should return nil")
	}
}

func TestLinkedList_Remove_index_greater_than_length(t *testing.T) {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)

	r := ll.Remove(4)

	if r != nil {
		t.Error("should return nil")
	}
}

func ExampleLinkedList_Reverse() {
	ll := New(10)
	ll.Push(11)
	ll.Push(12)
	ll.Push(13)
	ll.Push(14)
	ll.Push(15)

	ll.Reverse()

	ll.Print()

	// Output:
	// 15->14->13->12->11->10->nil
	// Head: 15
	// Tail: 10
	// Length: 6
}

func ExampleLinkedList_Reverse_two_items() {
	ll := New(10)
	ll.Push(11)

	ll.Reverse()

	ll.Print()

	// Output:
	// 11->10->nil
	// Head: 11
	// Tail: 10
	// Length: 2
}

func TestLinkedList_Reverse_empty_list(t *testing.T) {
	ll := LinkedList{}

	ll.Reverse()

	if ll.Length != 0 {
		t.Error("length\r\nwant: 0\r\ngot:", ll.Length)
	}
}

func TestLinkedList_Reverse_one_item(t *testing.T) {
	ll := New(10)

	ll.Reverse()

	if ll.Length != 1 {
		t.Error("length\r\nwant: 1\r\ngot:", ll.Length)
	}

	if ll.Head.Value != 10 {
		t.Error("head value\r\nwant: 10\r\ngot:", ll.Head.Value)
	}

	if ll.Tail.Value != 10 {
		t.Error("tail value\r\nwant: 10\r\ngot:", ll.Tail.Value)
	}
}

package doublylinkedlist

import "testing"

func TestDoublyLinkedList_IsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  bool
	}{
		{"empty list", []int{}, true},
		{"single-element list", []int{1}, true},
		{"two-element palindrome list", []int{1, 1}, true},
		{"two-element non-palindrome list", []int{1, 2}, false},
		{"palindrome list of odd length", []int{1, 2, 1}, true},
		{"palindrome list of even length", []int{1, 2, 2, 1}, true},
		{"non-palindrome list of odd length", []int{1, 2, 3}, false},
		{"non-palindrome list of even length", []int{1, 2, 3, 1}, false},
		{"non-palindrome list with repeated elements", []int{1, 1, 2, 2}, false},
		{"palindrome list with multiple same elements", []int{1, 1, 1, 1, 1, 1}, true},
	}

	for _, test := range tests {
		dll := FromValues(test.items...)

		r := dll.IsPalindrome()

		if r != test.want {
			t.Errorf("test \"%s\" failed\nwant: %v\ngot: %v", test.name, test.want, r)
		}
	}
}

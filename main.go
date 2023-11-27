// Package main won't be used much. It's just for the sake of having some examples running on the main function.
package main

import (
	"fmt"

	"github.com/ericmilaneze/data-structures-and-algorithms-in-go/linkedlist"
)

func main() {
	ll := linkedlist.New(9)
	ll.Print()

	fmt.Println()

	ll.Push(10)
	ll.Push(11)
	ll.Push(12)
	ll.Print()

	middleNode := ll.FindMiddleNode()
	fmt.Println(middleNode.Value)
}

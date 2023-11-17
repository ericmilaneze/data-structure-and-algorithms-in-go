package main

import (
	"fmt"

	"github.com/ericmilaneze/data-structure-and-algorithms-in-go/linkedlist"
)

func main() {
	ll := linkedlist.New(9)
	ll.Print()

	fmt.Println()

	ll.Push(10)
	ll.Push(11)
	ll.Push(12)
	ll.Print()
}

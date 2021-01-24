package main

import (
	"./lib"
)

func main() {
	linkedlist := lib.NewLinkedList([]int{1, 18, 4, 8, 123})
	linkedlist.PrintLinkedList()
	linkedlist.RemoveLast()
	linkedlist.PrintLinkedList()
	linkedlist.Append(321)
	linkedlist.PrintLinkedList()
}

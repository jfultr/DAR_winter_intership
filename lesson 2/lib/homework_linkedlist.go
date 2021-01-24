package lib

import "fmt"

// Node in LinkedList
type Node struct {
	Val  int
	Next *Node
}

// NewNode Func that create new node
func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

// LinkedList struct
type LinkedList struct {
	Head   *Node
	Length int
}

// NewLinkedList Func that create new node
func NewLinkedList(vals []int) *LinkedList {
	first := NewNode(vals[0])
	ll := &LinkedList{
		Head:   first,
		Length: 1,
	}
	for _, v := range vals[1:] {
		ll.Append(v)
	}
	return ll
}

// Append - appends new node to ll-s tail
func (ll *LinkedList) Append(val int) {
	node := NewNode(val)
	if ll.Head == nil {
		ll.Head = node
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	ll.Length++
	return
}

// RemoveLast remove last node
func (ll *LinkedList) RemoveLast() error {
	if ll.Head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}
	var prev *Node
	current := ll.Head
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	if prev != nil {
		prev.Next = nil
	} else {
		ll.Head = nil
	}
	ll.Length--
	return nil
}

// PrintLinkedList all nodes
func (ll LinkedList) PrintLinkedList() {
	toPrint := ll.Head
	for ll.Length != 0 {
		fmt.Printf("%d ", toPrint.Val)
		toPrint = toPrint.Next
		ll.Length--
	}
	fmt.Println("")
}

// struct linked list
// delete add print

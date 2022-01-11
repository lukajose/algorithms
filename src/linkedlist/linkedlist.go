package linkedlist

import (
	iterator "collections/interfaces"
	"fmt"
)

type LinkedList struct {
	size int
	root *Node
}

type LinkedListIterator struct {
	node *Node
}

// A comparable item that has to implement the compare to function
// for the linkedlist to know where to insert and make comparisons
type ComparableItem interface {
	// CompareTo should return
	// -1 if less than item
	// 1 if more than item
	// 0 if equal to item
	CompareTo(item interface{}) int
}

type Node struct {
	next *Node
	data ComparableItem
}

func NewNode(data ComparableItem) *Node {
	return &Node{
		next: nil,
		data: data,
	}
}

func NewList(data ComparableItem) *LinkedList {
	return &LinkedList{size: 1, root: NewNode(data)}
}

func (l *LinkedList) Print() {
	var curr = l.root

	for curr != nil {
		fmt.Printf("[%v]->", curr.data)
		curr = curr.next
	}
	fmt.Printf(" nil")

}

func (l *LinkedList) Append(data ComparableItem) {
	var curr = l.root
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = NewNode(data)
	l.size++
}

func (l *LinkedList) Insert(data ComparableItem) {
	var curr = l.root
	node := NewNode(data)
	node.next = curr
	l.root = node
	l.size++
}

func (l *LinkedList) InsertSorted(data ComparableItem) {
	// var curr = l.root

}

func (l *LinkedList) Search(data interface{}) interface{} {
	var curr = l.root
	for curr.next != nil {
		if curr.data.CompareTo(data) == 0 {
			return data
		}
		curr = curr.next
	}
	return nil
}

func (l *LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) Iterator() iterator.Iterator {
	return &LinkedListIterator{node: l.root}
}

func (l *LinkedListIterator) HasNext() bool {
	return l.node.next != nil
}

func (l *LinkedListIterator) Next() interface{} {
	var curr = l.node
	if l.HasNext() {
		l.node = l.node.next
	}
	return curr
}

package linkedlist

import (
	"collections/interfaces"
	"fmt"
)

type LinkedList struct {
	size int
	root *Node
}

type LinkedListIterator struct {
	node *Node
}

type Node struct {
	next *Node
	data interfaces.ComparableItem
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}

func NewNode(data interfaces.ComparableItem) *Node {
	return &Node{
		next: nil,
		data: data,
	}
}

func NewList() *LinkedList {
	return &LinkedList{size: 1, root: nil}
}

func (l *LinkedList) Print() {
	var curr = l.root

	for curr != nil {
		fmt.Printf("[%s]->", curr.data.Print())
		curr = curr.next
	}
	fmt.Printf(" nil")

}

func (l *LinkedList) Append(data interfaces.ComparableItem) {
	if l.root == nil {
		l.root = NewNode(data)
		return
	}
	var curr = l.root
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = NewNode(data)
	l.size++
}

func (l *LinkedList) Insert(data interfaces.ComparableItem) {
	if l.root == nil {
		l.root = NewNode(data)
		return
	}
	var curr = l.root
	node := NewNode(data)
	node.next = curr
	l.root = node
	l.size++
}

func (l *LinkedList) InsertSorted(data interfaces.ComparableItem) {
	if l.root == nil {
		l.root = NewNode(data)
		return
	}
	var curr = l.root
	for curr.next != nil {
		if data.CompareTo(curr) <= 0 && data.CompareTo(curr.next) == 1 {
			var temp = curr.next
			node := NewNode(data)
			curr.next = node
			node.next = temp
			l.size++
			return
		}
		curr = curr.next
	}
	curr.next = NewNode(data)
	l.size++

}

func (l *LinkedList) Search(data interface{}) interface{} {
	if l.root == nil {
		return nil
	}
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

func (l *LinkedList) Iterator() interfaces.Iterator {
	return &LinkedListIterator{node: l.root}
}

func (l *LinkedListIterator) HasNext() bool {
	if l.node == nil {
		return false
	}
	return l.node.next != nil
}

func (l *LinkedListIterator) Next() interface{} {
	var curr = l.node
	if l.HasNext() {
		l.node = l.node.next
	}
	return curr
}

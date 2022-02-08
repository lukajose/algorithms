package linkedlist

import (
	"fmt"
	"math/rand"
	"testing"
)

type IntItem struct {
	data int
}

func (i *IntItem) CompareTo(item interface{}) int {
	switch v := item.(type) {
	case int:
		if v < i.data {
			return -1
		}
		if v == i.data {
			return 0
		}
		return 1
	case *Node:
		if IntItem, ok := v.data.(*IntItem); ok {
			if IntItem.data < i.data {
				return -1
			}
			if IntItem.data == i.data {
				return 0
			}
			return 1
		}
		break
	}
	fmt.Println("item is:", item)
	panic(fmt.Errorf("Invalid type cannot compare int with item interface type"))
}

func (i *IntItem) Print() string {
	return fmt.Sprintf("%d", i.data)
}

func NewIntItem(data int) *IntItem {
	return &IntItem{data: data}
}
func TestInsert(t *testing.T) {
	l := NewList()
	for i := 0; i < 5; i++ {
		l.Insert(NewIntItem(i))
	}
	if l.Len() != 5 {
		t.Errorf("Expected at least 5 items but got %d", l.Len())
	}
	it := l.Iterator()
	l.Print()
	i := 4
	for it.HasNext() {
		node := it.Next().(*Node)
		if node.data.CompareTo(i) != 0 {
			t.Errorf("Expected item to be equal but got %v expected %d", node.data, i)
		}
		i--
	}
}

func TestInsertSorted(t *testing.T) {
	l := NewList()
	for i := 0; i < 5; i++ {
		l.InsertSorted(NewIntItem(rand.Int() % 100))
	}
	l.Print()
}

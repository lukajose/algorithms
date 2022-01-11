package linkedlist

import (
	"fmt"
	"testing"
)

type IntItem struct {
	data int
}

func (i *IntItem) CompareTo(item interface{}) int {
	if intItem, ok := item.(int); ok {
		if intItem < i.data {
			return -1
		}
		if intItem == i.data {
			return 0
		}
		return 1
	}
	panic(fmt.Errorf("Invalid type cannot compare int with item interface type"))
}

func NewIntItem(data int) *IntItem {
	return &IntItem{data: data}
}
func TestInsert(t *testing.T) {
	l := NewList(NewIntItem(0))
	var el []int
	for i := 1; i < 5; i++ {
		l.Insert(NewIntItem(i))
		el = append(el, i)
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

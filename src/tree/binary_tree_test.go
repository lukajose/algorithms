package tree

import (
	"collections/interfaces"
	"fmt"
	"testing"
)

type IntItem struct {
	data int
}

func (i IntItem) CompareTo(item interface{}) int {
	switch v := item.(type) {
	case int:
		if v < i.data {
			return -1
		}
		if v == i.data {
			return 0
		}
		return 1
	case interfaces.ComparableItem:
		if IntItem, ok := v.(*IntItem); ok {
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

func TestInsertion(t *testing.T) {
	tree := NewBinaryTree()
	tree.InOrder()
	tree.Put(&IntItem{10})
	tree.InOrder()
	tree.Put(&IntItem{59})
	tree.Put(&IntItem{1})
	tree.Put(&IntItem{15})
	tree.InOrder()
}

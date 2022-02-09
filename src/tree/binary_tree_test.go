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
	}
	panic(fmt.Errorf("Invalid type cannot compare IntItem with item passed as interface"))
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

func TestSearch(t *testing.T) {
	tree := NewBinaryTree()
	found := tree.Search(&IntItem{20})
	if found {
		tree.InOrder()
		t.Errorf("Returns true on empty tree")
	}
	tree.Put(&IntItem{59})
	found = tree.Search(&IntItem{59})
	if !found {
		tree.InOrder()
		t.Errorf("Failed to find inserted item 59")
	}
	tree.Put(&IntItem{1})
	tree.Put(&IntItem{20})
	tree.Put(&IntItem{5})
	found = tree.Search(&IntItem{20})
	if !found {
		tree.InOrder()
		t.Errorf("Failed to find inserted item 20")
	}
}

func TestDelete(t *testing.T) {
	tree := NewBinaryTree()
	tree.Put(&IntItem{1})
	tree.Put(&IntItem{20})
	tree.Put(&IntItem{5})
	tree.Put(&IntItem{8})
	tree.InOrder()
	tree.Delete(&IntItem{8})
	tree.InOrder()
	// tree.Put(&IntItem{-5})
	tree.Delete(&IntItem{1})
	tree.InOrder()
}

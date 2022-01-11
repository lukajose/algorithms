package heap

import (
	"fmt"
	"math"
)

type Comparable interface {
	CompareTo(item interface{}) (int, error)
}

type Heap struct {
	items    []Comparable
	capacity int
	size     int
}

type NodeMax struct {
	item int
}

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (n *NodeMax) CompareTo(item interface{}) (int, error) {
	val, ok := item.(int)
	if ok {
		switch true {
		case n.item > val:
			return 1, nil
		case n.item < val:
			return -1, nil
		default:
			return 0, nil
		}
	}
	return -1, fmt.Errorf("Invalid type comparison")
}

/*
Creates a new heap and returns a pointer to the heap
*/
func NewHeap(capacity int) *Heap {
	return &Heap{
		items:    make([]Comparable, 0, capacity),
		capacity: capacity,
		size:     0,
	}
}

/*
resize the heap to double the size
*/
func (h *Heap) resize() {
	h.items = h.items[:h.capacity*2]
}

func (h *Heap) heapifyUp() {
	if h.size == 1 {
		return
	}
	cIndex := h.size - 1
	child := h.items[cIndex]
	for cIndex >= 0 {
		pIndex := int(math.Floor(float64(cIndex) / 2))
		greater, ok := child.CompareTo(h.items[pIndex])
		if ok != nil && greater > 0 {
			// swap if element child is greater than parent
			h.items[pIndex], h.items[cIndex] = h.items[cIndex], h.items[pIndex]
		}
		cIndex = int(math.Floor(float64(cIndex) / 2))
	}

}

func (h *Heap) getMaxChild(parent int) (int, error) {
	c1, c2 := (parent * 2), (parent*2)+1
	if c1 > h.size && c2 > h.size {
		return -1, fmt.Errorf("both childs are out of range")
	}
	if c1 > h.size && c2 < h.size {
		return c1, nil
	}
	if c1 < h.size && c2 > h.size {
		return c2, nil
	}
	greater, err := h.items[c1].CompareTo(h.items[c2])
	if err != nil {
		return -1, err
	}
	if greater > 0 {
		return c1, nil
	}
	return c2, nil

}

func (h *Heap) heapifyDown() {
	if h.size <= 1 {
		return
	}
	parent := 0
	for parent < h.size {
		// get max of child nodes
		c, ok := h.getMaxChild(parent)
		if ok != nil {
			break
		}
		// if parent is smaller than child swap
		lower, ok := h.items[parent].CompareTo(h.items[c])
		if lower < 0 && ok != nil {
			// swap items
			h.items[parent], h.items[c] = h.items[c], h.items[parent]
		}
		parent++

	}

}

func (h *Heap) Add(item Comparable) {
	if h.size == h.capacity {
		h.resize()
	}
	h.items[h.size] = item
	h.size++
	h.heapifyUp()

}

func (h *Heap) Pop() Comparable {
	if len(h.items) == 0 {
		return nil
	}
	lindex := h.size - 1
	last := h.items[lindex]
	h.items = h.items[:lindex]
	h.heapifyDown()
	h.size = max(h.size-1, 0)
	return last
}

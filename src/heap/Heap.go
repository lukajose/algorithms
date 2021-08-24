package heap

import "math"

type Comparable interface {
	CompareTo(interface{}) int
}

type Heap struct {
	items    []Comparable
	capacity int
	next     int
}

/*
Creates a new heap and returns a pointer to the heap
*/
func NewHeap(capacity int) *Heap {
	return &Heap{
		items:    make([]Comparable, 0, capacity),
		capacity: capacity,
	}
}

/*
resize the heap to double the size
*/
func (h *Heap) resize() {
	h.items = h.items[:h.capacity*2]
}

func (h *Heap) heapifyUp() {
	cIndex := len(h.items) - 1
	child := h.items[cIndex]
	for cIndex >= 0 {
		pIndex := int(math.Floor(float64(cIndex) / 2))
		if child.CompareTo(h.items[pIndex]) > 0 {
			// swap if element child is greater than parent
			h.items[pIndex], h.items[cIndex] = h.items[cIndex], h.items[pIndex]
		}
		cIndex = int(math.Floor(float64(cIndex) / 2))

	}

}

func (h *Heap) heapifyDown() {

}

func (h *Heap) Add(item Comparable) {
	if h.next > h.capacity {
		h.resize()
	}
	h.items[len(h.items)-1] = item
	h.heapifyUp()
}

func (h *Heap) Pop() Comparable {
	if len(h.items) == 0 {
		return nil
	}
	lindex := len(h.items) - 1
	last := h.items[lindex]
	h.items = h.items[:lindex]
	h.heapifyDown()
	return last
}

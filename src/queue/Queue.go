package queue

import (
	iter "collections/interfaces"
)

type Queue struct {
	arr []interface{}
}

func NewQueue(N int) *Queue {
	var Queue = &Queue{
		arr: make([]interface{}, N),
	}
	return Queue
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) Push(el interface{}) {
	q.arr = append(q.arr, el)
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	index := 0
	el := q.arr[index]
	q.arr = q.arr[:index]
	return el, true
}

func (q *Queue) Iterator() iter.Iterator {
	return &QueueIterator{
		index: 0,
		arr:   q.arr,
	}
}

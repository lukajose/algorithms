package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue(3)
	it := queue.Iterator()
	queue.Push(1)
	queue.Push("hey there new string")
	queue.Push(1.143)
	for it.HasNext() {
		fmt.Sprintf("%v", it.Next())
	}

}

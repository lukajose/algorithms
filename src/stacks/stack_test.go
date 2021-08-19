package stacks

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(3)
	stack.Push(1)
	if stack.Size() < 1 {
		t.Errorf("was expecting at least one element")
	}
	stack.Push("this is another element")
	if stack.Size() < 2 {
		t.Errorf("was expecting at least one element")
	}

}

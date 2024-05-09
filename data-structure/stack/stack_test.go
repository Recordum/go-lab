package stack

import "testing"


// 
func TestPush(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	
	if stack.Len() != 6 {
		t.Errorf("Expected 6 but got %d", stack.Len())
	}
}
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
//가장 최근에 추가된 요소를 제거하고 반환한다.
func TestPop(t *testing.T){
	stack := New()
	stack.Push(1)
	stack.Push(2)
	popValue := stack.Pop()
	
	if popValue != 2 {
		t.Errorf("Expected 2 but got %d", popValue)
	}

	if stack.Len() != 1 {
		t.Errorf("Expected 1 but got %d", stack.Len())
	}
}

//stack이 비어있을 경우 pop을 호출하면 nil을 반환한다.
func TestPopEmpty(t *testing.T){
	stack := New()
	popValue := stack.Pop()
	
	if popValue != nil {
		t.Errorf("Expected nil but got %d", popValue)
	}
}
//stack이 비어있을 경우 true 반환 아닐경우 false 반환
func TestIsEmpty(t *testing.T) {
	stack := New()
	if !stack.IsEmpty() {
		t.Errorf("Expected true but got false")
	}
	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("Expected false but got true")
	}
}


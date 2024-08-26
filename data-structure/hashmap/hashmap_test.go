package hashmap

import (
	"testing"
)

func TestSet(t *testing.T) {
	newMap := New()
	newMap.Set("key1", 1)
	newMap.Set("key2", 2)
	newMap.Set("key3", 3)
	newMap.Set("key4", 4)
	newMap.Set("key5", 5)
	newMap.Set("key6", 6)
	if newMap.Size() != 6 {
		t.Errorf("Expected 6 but got %d", newMap.Size())
	}

	if newMap.Get("key6") != 6 {
		t.Errorf("Expected 6 but got %d", newMap.Get("key6"))
	}

	if newMap.Get("key7") != nil {
		t.Errorf("Expected nil but got %d", newMap.Get("key7"))
	}
}

// 중복된 key 값으로 Set을 호출할 경우 기존의 값을 덮어쓴다.
func TestSetDuplicateKey(t *testing.T) {
	newMap := New()
	newMap.Set("key1", 1)
	newMap.Set("key1", 2)

	if newMap.Size() != 1 {
		t.Errorf("Size Expected 1 but got %d", newMap.Size())
	}

	if newMap.Get("key1") != 2 {
		t.Errorf("Expected 2 but got %d", newMap.Get("key1"))
	}
}

func TestResize(t *testing.T) {
	newMap := New()
	for i := 0; i < 100; i++ {
		newMap.Set(i, i)
	}
	if newMap.Size() != 100 {
		t.Errorf("Expected 100 but got %d", newMap.Size())
	}
	if newMap.Capacity() != 128 {
		t.Errorf("Expected 128 but got %d", newMap.Capacity())
	}
}
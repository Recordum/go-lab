package btree

import (
	"testing"
)

func TestNewBTree(t *testing.T) {
	tests := []struct {
		name   string
		degree int
		panic  bool
	}{
		{"Valid degree 2", 2, false},
		{"Valid degree 3", 3, false},
		{"Valid degree 10", 10, false},
		{"Invalid degree 1", 1, true},
		{"Invalid degree 0", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.panic {
					t.Errorf("NewBTree() panic = %v, want %v", r != nil, tt.panic)
				}
			}()
			bt := NewBTree[int](tt.degree)
			if !tt.panic && bt == nil {
				t.Error("NewBTree() returned nil")
			}
		})
	}
}

func TestInsertAndSearch(t *testing.T) {
	bt := NewBTree[int](3)

	// Test insert and search
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, key := range keys {
		bt.Insert(key)
	}

	for _, key := range keys {
		if !bt.Search(key) {
			t.Errorf("Search(%d) = false, want true", key)
		}
	}

	// Test search for non-existent keys
	nonExistent := []int{1, 8, 15, 25, 100}
	for _, key := range nonExistent {
		if bt.Search(key) {
			t.Errorf("Search(%d) = true, want false", key)
		}
	}
}

func TestInsertDuplicates(t *testing.T) {
	bt := NewBTree[int](3)

	keys := []int{10, 20, 10, 20, 30}

	for _, key := range keys {
		bt.Insert(key)
	}

	traversal := bt.Traverse()
	// B-tree allows duplicates, so we should have 5 elements
	if len(traversal) != 5 {
		t.Errorf("Traverse() length = %d, want 5", len(traversal))
	}
}

func TestTraverse(t *testing.T) {
	bt := NewBTree[int](3)

	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, key := range keys {
		bt.Insert(key)
	}

	traversal := bt.Traverse()

	// Check if traversal is in sorted order
	for i := 1; i < len(traversal); i++ {
		if traversal[i] < traversal[i-1] {
			t.Errorf("Traverse() not sorted: %v", traversal)
			break
		}
	}

	// Check if all keys are present
	if len(traversal) != len(keys) {
		t.Errorf("Traverse() length = %d, want %d", len(traversal), len(keys))
	}
}

func TestDelete(t *testing.T) {
	bt := NewBTree[int](3)

	keys := []int{10, 20, 5, 6, 12, 30, 7, 17, 15, 25, 40, 50}

	for _, key := range keys {
		bt.Insert(key)
	}

	// Delete some keys
	deleteKeys := []int{6, 13, 7, 20}

	for _, key := range deleteKeys {
		result := bt.Delete(key)
		// 13 doesn't exist, so it should return false
		if key == 13 && result {
			t.Errorf("Delete(%d) = true, want false", key)
		}
		if key != 13 && !result {
			t.Errorf("Delete(%d) = false, want true", key)
		}
	}

	// Verify deleted keys are not found
	for _, key := range deleteKeys {
		if key != 13 && bt.Search(key) {
			t.Errorf("Search(%d) = true after deletion, want false", key)
		}
	}

	// Verify remaining keys are still found
	remainingKeys := []int{10, 5, 12, 30, 17, 15, 25, 40, 50}
	for _, key := range remainingKeys {
		if !bt.Search(key) {
			t.Errorf("Search(%d) = false, want true", key)
		}
	}
}

func TestDeleteFromEmptyTree(t *testing.T) {
	bt := NewBTree[int](3)

	if bt.Delete(10) {
		t.Error("Delete(10) from empty tree = true, want false")
	}
}

func TestStringBTree(t *testing.T) {
	bt := NewBTree[string](3)

	words := []string{"apple", "banana", "cherry", "date", "elderberry"}

	for _, word := range words {
		bt.Insert(word)
	}

	for _, word := range words {
		if !bt.Search(word) {
			t.Errorf("Search(%s) = false, want true", word)
		}
	}

	traversal := bt.Traverse()
	if len(traversal) != len(words) {
		t.Errorf("Traverse() length = %d, want %d", len(traversal), len(words))
	}
}

func TestHeight(t *testing.T) {
	bt := NewBTree[int](3)

	// Empty tree height should be 1 (root only)
	height := bt.Height()
	if height != 1 {
		t.Errorf("Height() = %d, want 1", height)
	}

	// Add some keys
	for i := 1; i <= 10; i++ {
		bt.Insert(i)
	}

	height = bt.Height()
	if height < 1 {
		t.Errorf("Height() = %d, should be at least 1", height)
	}
}

func TestLargeDataSet(t *testing.T) {
	bt := NewBTree[int](5)

	// Insert 1000 keys
	n := 1000
	for i := 0; i < n; i++ {
		bt.Insert(i)
	}

	// Verify all keys exist
	for i := 0; i < n; i++ {
		if !bt.Search(i) {
			t.Errorf("Search(%d) = false, want true", i)
		}
	}

	// Delete half the keys
	for i := 0; i < n; i += 2 {
		if !bt.Delete(i) {
			t.Errorf("Delete(%d) = false, want true", i)
		}
	}

	// Verify deleted keys are gone
	for i := 0; i < n; i += 2 {
		if bt.Search(i) {
			t.Errorf("Search(%d) = true after deletion, want false", i)
		}
	}

	// Verify remaining keys still exist
	for i := 1; i < n; i += 2 {
		if !bt.Search(i) {
			t.Errorf("Search(%d) = false, want true", i)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	bt := NewBTree[int](5)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bt.Insert(i)
	}
}

func BenchmarkSearch(b *testing.B) {
	bt := NewBTree[int](5)
	n := 10000

	for i := 0; i < n; i++ {
		bt.Insert(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bt.Search(i % n)
	}
}

func BenchmarkDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bt := NewBTree[int](5)
		n := 1000
		for j := 0; j < n; j++ {
			bt.Insert(j)
		}
		b.StartTimer()

		for j := 0; j < n; j++ {
			bt.Delete(j)
		}
	}
}

package btree

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		degree int
		want   int
	}{
		{"Valid degree", 3, 3},
		{"Minimum degree", 2, 2},
		{"Invalid degree - too small", 1, 2},
		{"Invalid degree - zero", 0, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := New(tt.degree)
			if bt.degree != tt.want {
				t.Errorf("New(%d).degree = %d, want %d", tt.degree, bt.degree, tt.want)
			}
			if bt.root == nil {
				t.Error("New() root is nil")
			}
			if !bt.root.isLeaf {
				t.Error("New() root should be a leaf")
			}
		})
	}
}

func TestInsertAndSearch(t *testing.T) {
	bt := New(3)
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}

	// Insert keys
	for _, key := range keys {
		bt.Insert(key)
	}

	// Search for inserted keys
	for _, key := range keys {
		if !bt.Search(key) {
			t.Errorf("Search(%d) = false, want true", key)
		}
	}

	// Search for non-existent keys
	nonExistent := []int{1, 15, 25, 100}
	for _, key := range nonExistent {
		if bt.Search(key) {
			t.Errorf("Search(%d) = true, want false", key)
		}
	}
}

func TestInsertDuplicates(t *testing.T) {
	bt := New(3)
	keys := []int{10, 20, 10, 30, 20}

	for _, key := range keys {
		bt.Insert(key)
	}

	traversed := bt.Traverse()
	expected := []int{10, 10, 20, 20, 30}

	if len(traversed) != len(expected) {
		t.Errorf("Traverse() length = %d, want %d", len(traversed), len(expected))
	}

	for i := range expected {
		if traversed[i] != expected[i] {
			t.Errorf("Traverse()[%d] = %d, want %d", i, traversed[i], expected[i])
		}
	}
}

func TestDelete(t *testing.T) {
	bt := New(3)
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}

	// Insert keys
	for _, key := range keys {
		bt.Insert(key)
	}

	// Delete some keys
	deleteKeys := []int{6, 12, 30}
	for _, key := range deleteKeys {
		bt.Delete(key)
		if bt.Search(key) {
			t.Errorf("After Delete(%d), Search(%d) = true, want false", key, key)
		}
	}

	// Verify remaining keys
	remainingKeys := []int{10, 20, 5, 7, 17}
	for _, key := range remainingKeys {
		if !bt.Search(key) {
			t.Errorf("Search(%d) = false, want true", key)
		}
	}
}

func TestDeleteFromLeaf(t *testing.T) {
	bt := New(3)
	keys := []int{1, 2, 3, 4, 5, 6, 7}

	for _, key := range keys {
		bt.Insert(key)
	}

	bt.Delete(4)

	if bt.Search(4) {
		t.Error("Search(4) = true after deletion, want false")
	}

	traversed := bt.Traverse()
	expected := []int{1, 2, 3, 5, 6, 7}

	if len(traversed) != len(expected) {
		t.Errorf("Traverse() length = %d, want %d", len(traversed), len(expected))
	}

	for i := range expected {
		if traversed[i] != expected[i] {
			t.Errorf("Traverse()[%d] = %d, want %d", i, traversed[i], expected[i])
		}
	}
}

func TestTraverse(t *testing.T) {
	bt := New(3)
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, key := range keys {
		bt.Insert(key)
	}

	traversed := bt.Traverse()
	expected := []int{5, 6, 7, 10, 12, 17, 20, 30}

	if len(traversed) != len(expected) {
		t.Errorf("Traverse() length = %d, want %d", len(traversed), len(expected))
	}

	for i := range expected {
		if traversed[i] != expected[i] {
			t.Errorf("Traverse()[%d] = %d, want %d", i, traversed[i], expected[i])
		}
	}
}

func TestLargeInsertions(t *testing.T) {
	bt := New(5)
	n := 1000

	// Insert keys from 1 to n
	for i := 1; i <= n; i++ {
		bt.Insert(i)
	}

	// Verify all keys exist
	for i := 1; i <= n; i++ {
		if !bt.Search(i) {
			t.Errorf("Search(%d) = false, want true", i)
		}
	}

	// Verify traverse returns sorted order
	traversed := bt.Traverse()
	if len(traversed) != n {
		t.Errorf("Traverse() length = %d, want %d", len(traversed), n)
	}

	for i := 0; i < n; i++ {
		if traversed[i] != i+1 {
			t.Errorf("Traverse()[%d] = %d, want %d", i, traversed[i], i+1)
			break
		}
	}
}

func TestEmptyTree(t *testing.T) {
	bt := New(3)

	if bt.Search(10) {
		t.Error("Search(10) on empty tree = true, want false")
	}

	traversed := bt.Traverse()
	if len(traversed) != 0 {
		t.Errorf("Traverse() on empty tree length = %d, want 0", len(traversed))
	}
}

func TestSingleElement(t *testing.T) {
	bt := New(3)
	bt.Insert(42)

	if !bt.Search(42) {
		t.Error("Search(42) = false, want true")
	}

	traversed := bt.Traverse()
	if len(traversed) != 1 || traversed[0] != 42 {
		t.Errorf("Traverse() = %v, want [42]", traversed)
	}

	bt.Delete(42)
	if bt.Search(42) {
		t.Error("Search(42) after deletion = true, want false")
	}
}

func BenchmarkInsert(b *testing.B) {
	bt := New(5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bt.Insert(i)
	}
}

func BenchmarkSearch(b *testing.B) {
	bt := New(5)
	for i := 0; i < 10000; i++ {
		bt.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bt.Search(i % 10000)
	}
}

func BenchmarkDelete(b *testing.B) {
	bt := New(5)
	for i := 0; i < b.N*2; i++ {
		bt.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bt.Delete(i)
	}
}

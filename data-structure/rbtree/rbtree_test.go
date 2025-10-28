package rbtree

import (
	"testing"
)

func TestNewRBTree(t *testing.T) {
	tree := NewRBTree()
	if tree == nil {
		t.Fatal("Expected non-nil tree")
	}
	if tree.Size != 0 {
		t.Errorf("Expected size 0, got %d", tree.Size)
	}
	if tree.Root != nil {
		t.Error("Expected nil root")
	}
}

func TestInsertAndSearch(t *testing.T) {
	tree := NewRBTree()

	// Insert single element
	tree.Insert(10, "ten")
	if tree.Size != 1 {
		t.Errorf("Expected size 1, got %d", tree.Size)
	}

	// Search for existing element
	value, found := tree.Search(10)
	if !found {
		t.Error("Expected to find key 10")
	}
	if value != "ten" {
		t.Errorf("Expected value 'ten', got %v", value)
	}

	// Search for non-existing element
	_, found = tree.Search(20)
	if found {
		t.Error("Expected not to find key 20")
	}
}

func TestInsertMultiple(t *testing.T) {
	tree := NewRBTree()
	keys := []int{10, 20, 30, 15, 25, 5, 1}

	for _, key := range keys {
		tree.Insert(key, key*10)
	}

	if tree.Size != len(keys) {
		t.Errorf("Expected size %d, got %d", len(keys), tree.Size)
	}

	// Verify all keys can be found
	for _, key := range keys {
		value, found := tree.Search(key)
		if !found {
			t.Errorf("Expected to find key %d", key)
		}
		if value != key*10 {
			t.Errorf("Expected value %d, got %v", key*10, value)
		}
	}

	// Verify root is black
	if tree.Root.Color != Black {
		t.Error("Root should be black")
	}
}

func TestUpdateValue(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(10, "initial")
	tree.Insert(10, "updated")

	if tree.Size != 1 {
		t.Errorf("Expected size 1 after update, got %d", tree.Size)
	}

	value, _ := tree.Search(10)
	if value != "updated" {
		t.Errorf("Expected value 'updated', got %v", value)
	}
}

func TestDelete(t *testing.T) {
	tree := NewRBTree()
	keys := []int{10, 20, 30, 15, 25, 5, 1}

	for _, key := range keys {
		tree.Insert(key, key*10)
	}

	// Delete existing key
	deleted := tree.Delete(15)
	if !deleted {
		t.Error("Expected successful deletion")
	}
	if tree.Size != len(keys)-1 {
		t.Errorf("Expected size %d, got %d", len(keys)-1, tree.Size)
	}

	// Verify key is gone
	_, found := tree.Search(15)
	if found {
		t.Error("Key should not be found after deletion")
	}

	// Delete non-existing key
	deleted = tree.Delete(100)
	if deleted {
		t.Error("Expected failed deletion for non-existing key")
	}
}

func TestDeleteAll(t *testing.T) {
	tree := NewRBTree()
	keys := []int{10, 20, 30}

	for _, key := range keys {
		tree.Insert(key, key)
	}

	// Delete all keys
	for _, key := range keys {
		tree.Delete(key)
	}

	if tree.Size != 0 {
		t.Errorf("Expected size 0, got %d", tree.Size)
	}

	if tree.Root != nil {
		t.Error("Root should be nil after deleting all nodes")
	}
}

func TestInOrderTraversal(t *testing.T) {
	tree := NewRBTree()
	keys := []int{10, 5, 20, 3, 7, 15, 30}

	for _, key := range keys {
		tree.Insert(key, key)
	}

	var result []int
	tree.InOrderTraversal(func(key int, value interface{}) {
		result = append(result, key)
	})

	// In-order traversal should return sorted keys
	expected := []int{3, 5, 7, 10, 15, 20, 30}
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tree := NewRBTree()

	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}

	tree.Insert(10, "ten")
	if tree.IsEmpty() {
		t.Error("Tree should not be empty after insert")
	}

	tree.Delete(10)
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after deleting all elements")
	}
}

func TestClear(t *testing.T) {
	tree := NewRBTree()
	for i := 0; i < 10; i++ {
		tree.Insert(i, i)
	}

	tree.Clear()

	if !tree.IsEmpty() {
		t.Error("Tree should be empty after clear")
	}
	if tree.Root != nil {
		t.Error("Root should be nil after clear")
	}
	if tree.Size != 0 {
		t.Errorf("Size should be 0 after clear, got %d", tree.Size)
	}
}

func TestRotations(t *testing.T) {
	tree := NewRBTree()

	// Insert in ascending order to trigger rotations
	for i := 1; i <= 7; i++ {
		tree.Insert(i, i)
	}

	// Verify tree is balanced (root should not be 1 or 7)
	if tree.Root.Key == 1 || tree.Root.Key == 7 {
		t.Error("Tree appears unbalanced after insertions")
	}

	// Verify root is black
	if tree.Root.Color != Black {
		t.Error("Root should always be black")
	}
}

func TestLargeInsertions(t *testing.T) {
	tree := NewRBTree()
	n := 1000

	// Insert many elements
	for i := 0; i < n; i++ {
		tree.Insert(i, i*2)
	}

	if tree.Size != n {
		t.Errorf("Expected size %d, got %d", n, tree.Size)
	}

	// Verify all can be found
	for i := 0; i < n; i++ {
		value, found := tree.Search(i)
		if !found {
			t.Errorf("Expected to find key %d", i)
		}
		if value != i*2 {
			t.Errorf("Expected value %d, got %v", i*2, value)
		}
	}

	// Verify root is black
	if tree.Root.Color != Black {
		t.Error("Root should be black")
	}
}

func BenchmarkInsert(b *testing.B) {
	tree := NewRBTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(i, i)
	}
}

func BenchmarkSearch(b *testing.B) {
	tree := NewRBTree()
	n := 10000
	for i := 0; i < n; i++ {
		tree.Insert(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(i % n)
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		tree := NewRBTree()
		for j := 0; j < 1000; j++ {
			tree.Insert(j, j)
		}
		b.StartTimer()
		tree.Delete(500)
		b.StopTimer()
	}
}

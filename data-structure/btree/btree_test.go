package btree

import (
	"testing"
)

func TestNew(t *testing.T) {
	tree := New(3)
	if tree == nil {
		t.Error("Failed to create B-tree")
	}
	if tree.degree != 3 {
		t.Errorf("Expected degree 3, got %d", tree.degree)
	}
	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}
}

func TestNewPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for degree < 2")
		}
	}()
	New(1)
}

func TestInsertAndSearch(t *testing.T) {
	tree := New(3)

	// Insert some values
	tree.Insert(10, "ten")
	tree.Insert(20, "twenty")
	tree.Insert(5, "five")
	tree.Insert(15, "fifteen")

	// Search for inserted values
	if val, found := tree.Search(10); !found || val != "ten" {
		t.Errorf("Expected to find 10 with value 'ten', got %v, %v", val, found)
	}

	if val, found := tree.Search(20); !found || val != "twenty" {
		t.Errorf("Expected to find 20 with value 'twenty', got %v, %v", val, found)
	}

	if val, found := tree.Search(5); !found || val != "five" {
		t.Errorf("Expected to find 5 with value 'five', got %v, %v", val, found)
	}

	// Search for non-existent key
	if _, found := tree.Search(99); found {
		t.Error("Should not find non-existent key 99")
	}
}

func TestInsertMany(t *testing.T) {
	tree := New(3)

	// Insert many values to trigger splits
	for i := 1; i <= 20; i++ {
		tree.Insert(i, i*10)
	}

	// Verify all values
	for i := 1; i <= 20; i++ {
		val, found := tree.Search(i)
		if !found {
			t.Errorf("Expected to find key %d", i)
		}
		if val.(int) != i*10 {
			t.Errorf("Expected value %d for key %d, got %v", i*10, i, val)
		}
	}

	if tree.Size() != 20 {
		t.Errorf("Expected size 20, got %d", tree.Size())
	}
}

func TestDelete(t *testing.T) {
	tree := New(3)

	// Insert values
	for i := 1; i <= 10; i++ {
		tree.Insert(i, i*10)
	}

	// Delete a value
	if !tree.Delete(5) {
		t.Error("Failed to delete key 5")
	}

	// Verify deletion
	if _, found := tree.Search(5); found {
		t.Error("Key 5 should have been deleted")
	}

	// Verify other values still exist
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		if _, found := tree.Search(i); !found {
			t.Errorf("Key %d should still exist", i)
		}
	}
}

func TestDeleteMany(t *testing.T) {
	tree := New(3)

	// Insert many values
	for i := 1; i <= 20; i++ {
		tree.Insert(i, i*10)
	}

	// Delete half of them
	for i := 1; i <= 20; i += 2 {
		if !tree.Delete(i) {
			t.Errorf("Failed to delete key %d", i)
		}
	}

	// Verify deletions
	for i := 1; i <= 20; i++ {
		_, found := tree.Search(i)
		if i%2 == 1 && found {
			t.Errorf("Key %d should have been deleted", i)
		}
		if i%2 == 0 && !found {
			t.Errorf("Key %d should still exist", i)
		}
	}

	if tree.Size() != 10 {
		t.Errorf("Expected size 10, got %d", tree.Size())
	}
}

func TestDeleteNonExistent(t *testing.T) {
	tree := New(3)
	tree.Insert(10, "ten")

	if tree.Delete(99) {
		t.Error("Should not be able to delete non-existent key")
	}
}

func TestHeight(t *testing.T) {
	tree := New(3)

	if tree.Height() != 1 {
		t.Errorf("Empty tree height should be 1, got %d", tree.Height())
	}

	// Insert values to increase height
	for i := 1; i <= 10; i++ {
		tree.Insert(i, i)
	}

	height := tree.Height()
	if height < 1 {
		t.Errorf("Height should be at least 1, got %d", height)
	}
}

func TestSize(t *testing.T) {
	tree := New(3)

	if tree.Size() != 0 {
		t.Errorf("Empty tree size should be 0, got %d", tree.Size())
	}

	for i := 1; i <= 15; i++ {
		tree.Insert(i, i)
		if tree.Size() != i {
			t.Errorf("After inserting %d items, size should be %d, got %d", i, i, tree.Size())
		}
	}

	for i := 1; i <= 5; i++ {
		tree.Delete(i)
		expected := 15 - i
		if tree.Size() != expected {
			t.Errorf("After deleting %d items, size should be %d, got %d", i, expected, tree.Size())
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tree := New(3)

	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}

	tree.Insert(10, "ten")
	if tree.IsEmpty() {
		t.Error("Tree with one element should not be empty")
	}

	tree.Delete(10)
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after deleting all elements")
	}
}

func TestDifferentDegrees(t *testing.T) {
	degrees := []int{2, 3, 4, 5, 10}

	for _, degree := range degrees {
		tree := New(degree)

		// Insert values
		for i := 1; i <= 50; i++ {
			tree.Insert(i, i*10)
		}

		// Verify all values
		for i := 1; i <= 50; i++ {
			val, found := tree.Search(i)
			if !found {
				t.Errorf("Degree %d: Expected to find key %d", degree, i)
			}
			if val.(int) != i*10 {
				t.Errorf("Degree %d: Expected value %d for key %d, got %v", degree, i*10, i, val)
			}
		}
	}
}

func TestComplexOperations(t *testing.T) {
	tree := New(3)

	// Mix of insertions and deletions
	for i := 1; i <= 20; i++ {
		tree.Insert(i, i*10)
	}

	for i := 5; i <= 10; i++ {
		tree.Delete(i)
	}

	for i := 21; i <= 30; i++ {
		tree.Insert(i, i*10)
	}

	for i := 15; i <= 20; i++ {
		tree.Delete(i)
	}

	// Verify remaining values
	for i := 1; i <= 4; i++ {
		if _, found := tree.Search(i); !found {
			t.Errorf("Key %d should exist", i)
		}
	}

	for i := 11; i <= 14; i++ {
		if _, found := tree.Search(i); !found {
			t.Errorf("Key %d should exist", i)
		}
	}

	for i := 21; i <= 30; i++ {
		if _, found := tree.Search(i); !found {
			t.Errorf("Key %d should exist", i)
		}
	}

	// Verify deleted values
	for i := 5; i <= 20; i++ {
		if i > 14 || (i >= 5 && i <= 10) {
			if _, found := tree.Search(i); found {
				t.Errorf("Key %d should have been deleted", i)
			}
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	tree := New(3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(i, i*10)
	}
}

func BenchmarkSearch(b *testing.B) {
	tree := New(3)
	for i := 0; i < 10000; i++ {
		tree.Insert(i, i*10)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(i % 10000)
	}
}

func BenchmarkDelete(b *testing.B) {
	tree := New(3)
	for i := 0; i < b.N*2; i++ {
		tree.Insert(i, i*10)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(i)
	}
}

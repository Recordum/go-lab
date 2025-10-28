package rbtree_test

import (
	"fmt"
	"github.com/Recordum/go-lab/data-structure/rbtree"
)

func ExampleRBTree_Insert() {
	tree := rbtree.NewRBTree()

	tree.Insert(10, "Apple")
	tree.Insert(20, "Banana")
	tree.Insert(5, "Cherry")

	fmt.Println("Tree size:", tree.GetSize())
	// Output: Tree size: 3
}

func ExampleRBTree_Search() {
	tree := rbtree.NewRBTree()

	tree.Insert(10, "Apple")
	tree.Insert(20, "Banana")
	tree.Insert(5, "Cherry")

	value, found := tree.Search(20)
	if found {
		fmt.Printf("Found: %v\n", value)
	}

	_, found = tree.Search(100)
	fmt.Printf("Key 100 found: %v\n", found)

	// Output:
	// Found: Banana
	// Key 100 found: false
}

func ExampleRBTree_Delete() {
	tree := rbtree.NewRBTree()

	tree.Insert(10, "Apple")
	tree.Insert(20, "Banana")
	tree.Insert(5, "Cherry")

	fmt.Println("Size before delete:", tree.GetSize())

	deleted := tree.Delete(20)
	fmt.Println("Deleted:", deleted)
	fmt.Println("Size after delete:", tree.GetSize())

	// Output:
	// Size before delete: 3
	// Deleted: true
	// Size after delete: 2
}

func ExampleRBTree_InOrderTraversal() {
	tree := rbtree.NewRBTree()

	tree.Insert(10, "Ten")
	tree.Insert(5, "Five")
	tree.Insert(20, "Twenty")
	tree.Insert(3, "Three")
	tree.Insert(7, "Seven")

	fmt.Println("In-order traversal:")
	tree.InOrderTraversal(func(key int, value interface{}) {
		fmt.Printf("%d:%v ", key, value)
	})

	// Output:
	// In-order traversal:
	// 3:Three 5:Five 7:Seven 10:Ten 20:Twenty
}

func ExampleRBTree_Clear() {
	tree := rbtree.NewRBTree()

	tree.Insert(10, "Apple")
	tree.Insert(20, "Banana")
	tree.Insert(5, "Cherry")

	fmt.Println("Size before clear:", tree.GetSize())
	fmt.Println("Is empty:", tree.IsEmpty())

	tree.Clear()

	fmt.Println("Size after clear:", tree.GetSize())
	fmt.Println("Is empty:", tree.IsEmpty())

	// Output:
	// Size before clear: 3
	// Is empty: false
	// Size after clear: 0
	// Is empty: true
}

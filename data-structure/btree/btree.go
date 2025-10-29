package btree

import (
	"fmt"
)

// Ordered is a constraint that permits any ordered type
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Node represents a node in the B-tree
type Node[T Ordered] struct {
	keys     []T
	children []*Node[T]
	isLeaf   bool
}

// BTree represents a B-tree data structure
type BTree[T Ordered] struct {
	root   *Node[T]
	degree int // minimum degree (minimum number of keys is degree-1)
}

// NewBTree creates a new B-tree with the given minimum degree
// degree must be at least 2
func NewBTree[T Ordered](degree int) *BTree[T] {
	if degree < 2 {
		panic("B-tree degree must be at least 2")
	}
	return &BTree[T]{
		root:   &Node[T]{isLeaf: true},
		degree: degree,
	}
}

// Search searches for a key in the B-tree
func (bt *BTree[T]) Search(key T) bool {
	return bt.searchNode(bt.root, key)
}

func (bt *BTree[T]) searchNode(node *Node[T], key T) bool {
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	if i < len(node.keys) && key == node.keys[i] {
		return true
	}

	if node.isLeaf {
		return false
	}

	return bt.searchNode(node.children[i], key)
}

// Insert inserts a key into the B-tree
func (bt *BTree[T]) Insert(key T) {
	root := bt.root

	// If root is full, split it
	if len(root.keys) == 2*bt.degree-1 {
		newRoot := &Node[T]{isLeaf: false}
		newRoot.children = append(newRoot.children, root)
		bt.splitChild(newRoot, 0)
		bt.root = newRoot
	}

	bt.insertNonFull(bt.root, key)
}

func (bt *BTree[T]) insertNonFull(node *Node[T], key T) {
	i := len(node.keys) - 1

	if node.isLeaf {
		// Insert key in sorted order
		node.keys = append(node.keys, key)
		for i >= 0 && key < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = key
	} else {
		// Find child to insert
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++

		// Split child if full
		if len(node.children[i].keys) == 2*bt.degree-1 {
			bt.splitChild(node, i)
			if key > node.keys[i] {
				i++
			}
		}
		bt.insertNonFull(node.children[i], key)
	}
}

func (bt *BTree[T]) splitChild(parent *Node[T], index int) {
	child := parent.children[index]
	mid := bt.degree - 1

	// Create new node for right half
	newChild := &Node[T]{isLeaf: child.isLeaf}
	newChild.keys = append([]T{}, child.keys[mid+1:]...)

	if !child.isLeaf {
		newChild.children = append([]*Node[T]{}, child.children[mid+1:]...)
		child.children = child.children[:mid+1]
	}

	// Move middle key up to parent
	parent.keys = append(parent.keys, child.keys[mid])
	copy(parent.keys[index+1:], parent.keys[index:])
	parent.keys[index] = child.keys[mid]

	// Insert new child
	parent.children = append(parent.children, nil)
	copy(parent.children[index+2:], parent.children[index+1:])
	parent.children[index+1] = newChild

	// Update original child
	child.keys = child.keys[:mid]
}

// Delete removes a key from the B-tree
func (bt *BTree[T]) Delete(key T) bool {
	if bt.root == nil {
		return false
	}

	deleted := bt.deleteFromNode(bt.root, key)

	// If root is empty after deletion, make its only child the new root
	if len(bt.root.keys) == 0 {
		if !bt.root.isLeaf && len(bt.root.children) > 0 {
			bt.root = bt.root.children[0]
		}
	}

	return deleted
}

func (bt *BTree[T]) deleteFromNode(node *Node[T], key T) bool {
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	if i < len(node.keys) && key == node.keys[i] {
		if node.isLeaf {
			// Case 1: Key in leaf node
			node.keys = append(node.keys[:i], node.keys[i+1:]...)
			return true
		}
		// Case 2: Key in internal node
		return bt.deleteFromInternal(node, i)
	}

	if node.isLeaf {
		return false
	}

	// Case 3: Key not in this node
	isInSubtree := i < len(node.keys)
	if len(node.children[i].keys) < bt.degree {
		bt.fill(node, i)
	}

	if isInSubtree && i > len(node.keys) {
		return bt.deleteFromNode(node.children[i-1], key)
	}
	return bt.deleteFromNode(node.children[i], key)
}

func (bt *BTree[T]) deleteFromInternal(node *Node[T], index int) bool {
	key := node.keys[index]

	if len(node.children[index].keys) >= bt.degree {
		// Get predecessor
		pred := bt.getPredecessor(node, index)
		node.keys[index] = pred
		return bt.deleteFromNode(node.children[index], pred)
	}

	if len(node.children[index+1].keys) >= bt.degree {
		// Get successor
		succ := bt.getSuccessor(node, index)
		node.keys[index] = succ
		return bt.deleteFromNode(node.children[index+1], succ)
	}

	// Merge with sibling
	bt.merge(node, index)
	return bt.deleteFromNode(node.children[index], key)
}

func (bt *BTree[T]) getPredecessor(node *Node[T], index int) T {
	curr := node.children[index]
	for !curr.isLeaf {
		curr = curr.children[len(curr.children)-1]
	}
	return curr.keys[len(curr.keys)-1]
}

func (bt *BTree[T]) getSuccessor(node *Node[T], index int) T {
	curr := node.children[index+1]
	for !curr.isLeaf {
		curr = curr.children[0]
	}
	return curr.keys[0]
}

func (bt *BTree[T]) fill(node *Node[T], index int) {
	// Borrow from previous sibling
	if index > 0 && len(node.children[index-1].keys) >= bt.degree {
		bt.borrowFromPrev(node, index)
		return
	}

	// Borrow from next sibling
	if index < len(node.children)-1 && len(node.children[index+1].keys) >= bt.degree {
		bt.borrowFromNext(node, index)
		return
	}

	// Merge with sibling
	if index < len(node.children)-1 {
		bt.merge(node, index)
	} else {
		bt.merge(node, index-1)
	}
}

func (bt *BTree[T]) borrowFromPrev(node *Node[T], index int) {
	child := node.children[index]
	sibling := node.children[index-1]

	// Move key from parent to child
	child.keys = append([]T{node.keys[index-1]}, child.keys...)

	// Move key from sibling to parent
	node.keys[index-1] = sibling.keys[len(sibling.keys)-1]
	sibling.keys = sibling.keys[:len(sibling.keys)-1]

	// Move child pointer
	if !child.isLeaf {
		child.children = append([]*Node[T]{sibling.children[len(sibling.children)-1]}, child.children...)
		sibling.children = sibling.children[:len(sibling.children)-1]
	}
}

func (bt *BTree[T]) borrowFromNext(node *Node[T], index int) {
	child := node.children[index]
	sibling := node.children[index+1]

	// Move key from parent to child
	child.keys = append(child.keys, node.keys[index])

	// Move key from sibling to parent
	node.keys[index] = sibling.keys[0]
	sibling.keys = sibling.keys[1:]

	// Move child pointer
	if !child.isLeaf {
		child.children = append(child.children, sibling.children[0])
		sibling.children = sibling.children[1:]
	}
}

func (bt *BTree[T]) merge(node *Node[T], index int) {
	child := node.children[index]
	sibling := node.children[index+1]

	// Pull key from current node and merge with right sibling
	child.keys = append(child.keys, node.keys[index])
	child.keys = append(child.keys, sibling.keys...)

	if !child.isLeaf {
		child.children = append(child.children, sibling.children...)
	}

	// Remove key from current node
	node.keys = append(node.keys[:index], node.keys[index+1:]...)

	// Remove child pointer
	node.children = append(node.children[:index+1], node.children[index+2:]...)
}

// Traverse performs an in-order traversal of the B-tree
func (bt *BTree[T]) Traverse() []T {
	var result []T
	bt.traverseNode(bt.root, &result)
	return result
}

func (bt *BTree[T]) traverseNode(node *Node[T], result *[]T) {
	i := 0
	for i < len(node.keys) {
		if !node.isLeaf {
			bt.traverseNode(node.children[i], result)
		}
		*result = append(*result, node.keys[i])
		i++
	}

	if !node.isLeaf {
		bt.traverseNode(node.children[i], result)
	}
}

// Height returns the height of the B-tree
func (bt *BTree[T]) Height() int {
	return bt.heightNode(bt.root)
}

func (bt *BTree[T]) heightNode(node *Node[T]) int {
	if node.isLeaf {
		return 1
	}
	return 1 + bt.heightNode(node.children[0])
}

// Print prints the B-tree structure (for debugging)
func (bt *BTree[T]) Print() {
	bt.printNode(bt.root, 0)
}

func (bt *BTree[T]) printNode(node *Node[T], level int) {
	if node == nil {
		return
	}

	fmt.Printf("Level %d: %v\n", level, node.keys)

	if !node.isLeaf {
		for _, child := range node.children {
			bt.printNode(child, level+1)
		}
	}
}

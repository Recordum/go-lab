package btree

import (
	"fmt"
)

// BTree represents a B-tree data structure
type BTree struct {
	root   *Node
	degree int // minimum degree (minimum number of keys = degree-1)
}

// Node represents a node in the B-tree
type Node struct {
	keys     []int
	values   []interface{}
	children []*Node
	isLeaf   bool
}

// New creates a new B-tree with the specified degree
// degree must be at least 2
func New(degree int) *BTree {
	if degree < 2 {
		panic("B-tree degree must be at least 2")
	}
	return &BTree{
		root:   newNode(true),
		degree: degree,
	}
}

// newNode creates a new node
func newNode(isLeaf bool) *Node {
	return &Node{
		keys:     make([]int, 0),
		values:   make([]interface{}, 0),
		children: make([]*Node, 0),
		isLeaf:   isLeaf,
	}
}

// Search searches for a key in the B-tree and returns its value
func (t *BTree) Search(key int) (interface{}, bool) {
	return t.searchNode(t.root, key)
}

// searchNode recursively searches for a key in a node
func (t *BTree) searchNode(node *Node, key int) (interface{}, bool) {
	if node == nil {
		return nil, false
	}

	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	// Key found
	if i < len(node.keys) && key == node.keys[i] {
		return node.values[i], true
	}

	// If leaf node, key doesn't exist
	if node.isLeaf {
		return nil, false
	}

	// Recursively search in child
	return t.searchNode(node.children[i], key)
}

// Insert inserts a key-value pair into the B-tree
func (t *BTree) Insert(key int, value interface{}) {
	root := t.root

	// If root is full, split it
	if len(root.keys) == 2*t.degree-1 {
		newRoot := newNode(false)
		newRoot.children = append(newRoot.children, t.root)
		t.splitChild(newRoot, 0)
		t.root = newRoot
		root = newRoot
	}

	t.insertNonFull(root, key, value)
}

// insertNonFull inserts a key into a non-full node
func (t *BTree) insertNonFull(node *Node, key int, value interface{}) {
	i := len(node.keys) - 1

	if node.isLeaf {
		// Insert key in sorted order
		node.keys = append(node.keys, 0)
		node.values = append(node.values, nil)

		for i >= 0 && key < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			node.values[i+1] = node.values[i]
			i--
		}
		node.keys[i+1] = key
		node.values[i+1] = value
	} else {
		// Find child to insert into
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++

		// Split child if full
		if len(node.children[i].keys) == 2*t.degree-1 {
			t.splitChild(node, i)
			if key > node.keys[i] {
				i++
			}
		}
		t.insertNonFull(node.children[i], key, value)
	}
}

// splitChild splits a full child of a node
func (t *BTree) splitChild(parent *Node, index int) {
	degree := t.degree
	fullChild := parent.children[index]
	newChild := newNode(fullChild.isLeaf)

	mid := degree - 1

	// Move second half of keys and values to new node
	newChild.keys = append(newChild.keys, fullChild.keys[mid+1:]...)
	newChild.values = append(newChild.values, fullChild.values[mid+1:]...)

	// If not leaf, move children too
	if !fullChild.isLeaf {
		newChild.children = append(newChild.children, fullChild.children[mid+1:]...)
		fullChild.children = fullChild.children[:mid+1]
	}

	// Move middle key up to parent
	parent.keys = append(parent.keys[:index], append([]int{fullChild.keys[mid]}, parent.keys[index:]...)...)
	parent.values = append(parent.values[:index], append([]interface{}{fullChild.values[mid]}, parent.values[index:]...)...)
	parent.children = append(parent.children[:index+1], append([]*Node{newChild}, parent.children[index+1:]...)...)

	// Truncate original child
	fullChild.keys = fullChild.keys[:mid]
	fullChild.values = fullChild.values[:mid]
}

// Delete removes a key from the B-tree
func (t *BTree) Delete(key int) bool {
	if t.root == nil {
		return false
	}

	deleted := t.deleteFromNode(t.root, key)

	// If root is empty after deletion, make its only child the new root
	if len(t.root.keys) == 0 {
		if !t.root.isLeaf && len(t.root.children) > 0 {
			t.root = t.root.children[0]
		}
	}

	return deleted
}

// deleteFromNode deletes a key from a node
func (t *BTree) deleteFromNode(node *Node, key int) bool {
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	if i < len(node.keys) && key == node.keys[i] {
		// Key found in this node
		if node.isLeaf {
			// Simply remove from leaf
			node.keys = append(node.keys[:i], node.keys[i+1:]...)
			node.values = append(node.values[:i], node.values[i+1:]...)
			return true
		}
		// Key in internal node - replace with predecessor or successor
		return t.deleteInternalNode(node, key, i)
	}

	if node.isLeaf {
		// Key not found
		return false
	}

	// Key might be in subtree
	isInSubtree := i < len(node.keys)
	if len(node.children[i].keys) < t.degree {
		// Child has minimum keys, need to rebalance
		t.rebalance(node, i)
		// After rebalancing, key position might change
		for i < len(node.keys) && key > node.keys[i] {
			i++
		}
		if i < len(node.keys) && key == node.keys[i] {
			return t.deleteInternalNode(node, key, i)
		}
		if i >= len(node.children) {
			return false
		}
	}

	if isInSubtree && i > len(node.keys) {
		i--
	}

	return t.deleteFromNode(node.children[i], key)
}

// deleteInternalNode handles deletion from internal node
func (t *BTree) deleteInternalNode(node *Node, key int, index int) bool {
	if node.isLeaf {
		node.keys = append(node.keys[:index], node.keys[index+1:]...)
		node.values = append(node.values[:index], node.values[index+1:]...)
		return true
	}

	if len(node.children[index].keys) >= t.degree {
		// Get predecessor
		pred := t.getPredecessor(node, index)
		node.keys[index] = pred.key
		node.values[index] = pred.value
		t.deleteFromNode(node.children[index], pred.key)
		return true
	}

	if len(node.children[index+1].keys) >= t.degree {
		// Get successor
		succ := t.getSuccessor(node, index)
		node.keys[index] = succ.key
		node.values[index] = succ.value
		t.deleteFromNode(node.children[index+1], succ.key)
		return true
	}

	// Merge with sibling
	t.merge(node, index)
	return t.deleteFromNode(node.children[index], key)
}

// getPredecessor gets the predecessor key-value pair
func (t *BTree) getPredecessor(node *Node, index int) *keyValue {
	curr := node.children[index]
	for !curr.isLeaf {
		curr = curr.children[len(curr.children)-1]
	}
	n := len(curr.keys)
	return &keyValue{key: curr.keys[n-1], value: curr.values[n-1]}
}

// getSuccessor gets the successor key-value pair
func (t *BTree) getSuccessor(node *Node, index int) *keyValue {
	curr := node.children[index+1]
	for !curr.isLeaf {
		curr = curr.children[0]
	}
	return &keyValue{key: curr.keys[0], value: curr.values[0]}
}

type keyValue struct {
	key   int
	value interface{}
}

// rebalance ensures child has enough keys before deletion
func (t *BTree) rebalance(node *Node, index int) {
	// Try to borrow from left sibling
	if index > 0 && len(node.children[index-1].keys) >= t.degree {
		t.borrowFromLeft(node, index)
		return
	}

	// Try to borrow from right sibling
	if index < len(node.children)-1 && len(node.children[index+1].keys) >= t.degree {
		t.borrowFromRight(node, index)
		return
	}

	// Merge with sibling
	if index > 0 {
		t.merge(node, index-1)
	} else {
		t.merge(node, index)
	}
}

// borrowFromLeft borrows a key from left sibling
func (t *BTree) borrowFromLeft(node *Node, childIndex int) {
	child := node.children[childIndex]
	sibling := node.children[childIndex-1]

	// Move parent key down to child
	child.keys = append([]int{node.keys[childIndex-1]}, child.keys...)
	child.values = append([]interface{}{node.values[childIndex-1]}, child.values...)

	// Move sibling's last key up to parent
	node.keys[childIndex-1] = sibling.keys[len(sibling.keys)-1]
	node.values[childIndex-1] = sibling.values[len(sibling.values)-1]

	// Move sibling's last child to child
	if !child.isLeaf {
		child.children = append([]*Node{sibling.children[len(sibling.children)-1]}, child.children...)
		sibling.children = sibling.children[:len(sibling.children)-1]
	}

	// Remove last key from sibling
	sibling.keys = sibling.keys[:len(sibling.keys)-1]
	sibling.values = sibling.values[:len(sibling.values)-1]
}

// borrowFromRight borrows a key from right sibling
func (t *BTree) borrowFromRight(node *Node, childIndex int) {
	child := node.children[childIndex]
	sibling := node.children[childIndex+1]

	// Move parent key down to child
	child.keys = append(child.keys, node.keys[childIndex])
	child.values = append(child.values, node.values[childIndex])

	// Move sibling's first key up to parent
	node.keys[childIndex] = sibling.keys[0]
	node.values[childIndex] = sibling.values[0]

	// Move sibling's first child to child
	if !child.isLeaf {
		child.children = append(child.children, sibling.children[0])
		sibling.children = sibling.children[1:]
	}

	// Remove first key from sibling
	sibling.keys = sibling.keys[1:]
	sibling.values = sibling.values[1:]
}

// merge merges a child with its sibling
func (t *BTree) merge(node *Node, index int) {
	child := node.children[index]
	sibling := node.children[index+1]

	// Pull key from parent and merge with right sibling
	child.keys = append(child.keys, node.keys[index])
	child.values = append(child.values, node.values[index])
	child.keys = append(child.keys, sibling.keys...)
	child.values = append(child.values, sibling.values...)

	if !child.isLeaf {
		child.children = append(child.children, sibling.children...)
	}

	// Remove key from parent
	node.keys = append(node.keys[:index], node.keys[index+1:]...)
	node.values = append(node.values[:index], node.values[index+1:]...)
	node.children = append(node.children[:index+1], node.children[index+2:]...)
}

// Height returns the height of the B-tree
func (t *BTree) Height() int {
	if t.root == nil {
		return 0
	}
	return t.heightNode(t.root)
}

func (t *BTree) heightNode(node *Node) int {
	if node.isLeaf {
		return 1
	}
	return 1 + t.heightNode(node.children[0])
}

// Size returns the total number of keys in the B-tree
func (t *BTree) Size() int {
	if t.root == nil {
		return 0
	}
	return t.sizeNode(t.root)
}

func (t *BTree) sizeNode(node *Node) int {
	count := len(node.keys)
	if !node.isLeaf {
		for _, child := range node.children {
			count += t.sizeNode(child)
		}
	}
	return count
}

// IsEmpty returns true if the tree is empty
func (t *BTree) IsEmpty() bool {
	return t.root == nil || len(t.root.keys) == 0
}

// Print prints the B-tree structure (for debugging)
func (t *BTree) Print() {
	if t.root == nil {
		fmt.Println("Empty tree")
		return
	}
	t.printNode(t.root, 0)
}

func (t *BTree) printNode(node *Node, level int) {
	if node == nil {
		return
	}

	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	fmt.Printf("%sKeys: %v\n", indent, node.keys)
	if !node.isLeaf {
		for _, child := range node.children {
			t.printNode(child, level+1)
		}
	}
}

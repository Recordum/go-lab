package btree

import (
	"fmt"
)

// BTree represents a B-tree data structure
type BTree struct {
	root   *Node
	degree int // minimum degree (t)
}

// Node represents a node in the B-tree
type Node struct {
	keys     []int
	children []*Node
	isLeaf   bool
}

// New creates a new B-tree with the specified minimum degree
// The minimum degree t must be at least 2
func New(degree int) *BTree {
	if degree < 2 {
		degree = 2
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
		children: make([]*Node, 0),
		isLeaf:   isLeaf,
	}
}

// Search searches for a key in the B-tree
func (bt *BTree) Search(key int) bool {
	return bt.searchNode(bt.root, key)
}

// searchNode recursively searches for a key in a node
func (bt *BTree) searchNode(node *Node, key int) bool {
	if node == nil {
		return false
	}

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
func (bt *BTree) Insert(key int) {
	root := bt.root

	// If root is full, split it
	if len(root.keys) == 2*bt.degree-1 {
		newRoot := newNode(false)
		newRoot.children = append(newRoot.children, bt.root)
		bt.splitChild(newRoot, 0)
		bt.root = newRoot
	}

	bt.insertNonFull(bt.root, key)
}

// insertNonFull inserts a key into a non-full node
func (bt *BTree) insertNonFull(node *Node, key int) {
	i := len(node.keys) - 1

	if node.isLeaf {
		// Insert key in sorted order
		node.keys = append(node.keys, 0)
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

		// If child is full, split it
		if len(node.children[i].keys) == 2*bt.degree-1 {
			bt.splitChild(node, i)
			if key > node.keys[i] {
				i++
			}
		}
		bt.insertNonFull(node.children[i], key)
	}
}

// splitChild splits a full child of a node
func (bt *BTree) splitChild(parent *Node, index int) {
	t := bt.degree
	fullChild := parent.children[index]
	newChild := newNode(fullChild.isLeaf)

	// Copy the second half of keys to new child
	newChild.keys = make([]int, t-1)
	copy(newChild.keys, fullChild.keys[t:])

	// If not leaf, copy the second half of children
	if !fullChild.isLeaf {
		newChild.children = make([]*Node, t)
		copy(newChild.children, fullChild.children[t:])
		fullChild.children = fullChild.children[:t]
	}

	// Move middle key up to parent
	midKey := fullChild.keys[t-1]
	fullChild.keys = fullChild.keys[:t-1]

	// Insert middle key into parent
	parent.keys = append(parent.keys, 0)
	copy(parent.keys[index+1:], parent.keys[index:])
	parent.keys[index] = midKey

	// Insert new child into parent
	parent.children = append(parent.children, nil)
	copy(parent.children[index+2:], parent.children[index+1:])
	parent.children[index+1] = newChild
}

// Delete removes a key from the B-tree
func (bt *BTree) Delete(key int) {
	bt.deleteFromNode(bt.root, key)

	// If root is empty after deletion, make its only child the new root
	if len(bt.root.keys) == 0 {
		if !bt.root.isLeaf && len(bt.root.children) > 0 {
			bt.root = bt.root.children[0]
		}
	}
}

// deleteFromNode deletes a key from a node
func (bt *BTree) deleteFromNode(node *Node, key int) {
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	if i < len(node.keys) && key == node.keys[i] {
		if node.isLeaf {
			// Case 1: Key is in leaf node
			node.keys = append(node.keys[:i], node.keys[i+1:]...)
		} else {
			// Case 2: Key is in internal node
			bt.deleteFromInternalNode(node, i)
		}
	} else if !node.isLeaf {
		// Case 3: Key is in subtree
		isInLastChild := i == len(node.keys)
		if len(node.children[i].keys) < bt.degree {
			bt.fill(node, i)
		}

		if isInLastChild && i > len(node.keys) {
			bt.deleteFromNode(node.children[i-1], key)
		} else {
			bt.deleteFromNode(node.children[i], key)
		}
	}
}

// deleteFromInternalNode deletes a key from an internal node
func (bt *BTree) deleteFromInternalNode(node *Node, index int) {
	key := node.keys[index]

	if len(node.children[index].keys) >= bt.degree {
		// Get predecessor and replace
		pred := bt.getPredecessor(node, index)
		node.keys[index] = pred
		bt.deleteFromNode(node.children[index], pred)
	} else if len(node.children[index+1].keys) >= bt.degree {
		// Get successor and replace
		succ := bt.getSuccessor(node, index)
		node.keys[index] = succ
		bt.deleteFromNode(node.children[index+1], succ)
	} else {
		// Merge with sibling
		bt.merge(node, index)
		bt.deleteFromNode(node.children[index], key)
	}
}

// getPredecessor gets the predecessor key
func (bt *BTree) getPredecessor(node *Node, index int) int {
	curr := node.children[index]
	for !curr.isLeaf {
		curr = curr.children[len(curr.children)-1]
	}
	return curr.keys[len(curr.keys)-1]
}

// getSuccessor gets the successor key
func (bt *BTree) getSuccessor(node *Node, index int) int {
	curr := node.children[index+1]
	for !curr.isLeaf {
		curr = curr.children[0]
	}
	return curr.keys[0]
}

// fill ensures a child has at least t keys
func (bt *BTree) fill(node *Node, index int) {
	// If previous sibling has at least t keys, borrow from it
	if index != 0 && len(node.children[index-1].keys) >= bt.degree {
		bt.borrowFromPrev(node, index)
	} else if index != len(node.keys) && len(node.children[index+1].keys) >= bt.degree {
		// If next sibling has at least t keys, borrow from it
		bt.borrowFromNext(node, index)
	} else {
		// Merge with sibling
		if index != len(node.keys) {
			bt.merge(node, index)
		} else {
			bt.merge(node, index-1)
		}
	}
}

// borrowFromPrev borrows a key from the previous sibling
func (bt *BTree) borrowFromPrev(node *Node, childIndex int) {
	child := node.children[childIndex]
	sibling := node.children[childIndex-1]

	// Move a key from parent to child
	child.keys = append([]int{node.keys[childIndex-1]}, child.keys...)

	// Move a key from sibling to parent
	node.keys[childIndex-1] = sibling.keys[len(sibling.keys)-1]
	sibling.keys = sibling.keys[:len(sibling.keys)-1]

	// Move child pointer if not leaf
	if !child.isLeaf {
		child.children = append([]*Node{sibling.children[len(sibling.children)-1]}, child.children...)
		sibling.children = sibling.children[:len(sibling.children)-1]
	}
}

// borrowFromNext borrows a key from the next sibling
func (bt *BTree) borrowFromNext(node *Node, childIndex int) {
	child := node.children[childIndex]
	sibling := node.children[childIndex+1]

	// Move a key from parent to child
	child.keys = append(child.keys, node.keys[childIndex])

	// Move a key from sibling to parent
	node.keys[childIndex] = sibling.keys[0]
	sibling.keys = sibling.keys[1:]

	// Move child pointer if not leaf
	if !child.isLeaf {
		child.children = append(child.children, sibling.children[0])
		sibling.children = sibling.children[1:]
	}
}

// merge merges a child with its sibling
func (bt *BTree) merge(node *Node, index int) {
	child := node.children[index]
	sibling := node.children[index+1]

	// Pull key from current node and merge with right sibling
	child.keys = append(child.keys, node.keys[index])
	child.keys = append(child.keys, sibling.keys...)

	// Copy child pointers
	if !child.isLeaf {
		child.children = append(child.children, sibling.children...)
	}

	// Remove key from current node
	node.keys = append(node.keys[:index], node.keys[index+1:]...)

	// Remove child pointer
	node.children = append(node.children[:index+1], node.children[index+2:]...)
}

// Traverse performs an in-order traversal of the B-tree
func (bt *BTree) Traverse() []int {
	result := make([]int, 0)
	bt.traverseNode(bt.root, &result)
	return result
}

// traverseNode recursively traverses a node
func (bt *BTree) traverseNode(node *Node, result *[]int) {
	if node == nil {
		return
	}

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

// Print prints the B-tree structure
func (bt *BTree) Print() {
	bt.printNode(bt.root, 0)
}

// printNode recursively prints a node
func (bt *BTree) printNode(node *Node, level int) {
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

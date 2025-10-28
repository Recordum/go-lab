package rbtree

import "fmt"

// Color represents the color of a Red-Black Tree node
type Color bool

const (
	Red   Color = true
	Black Color = false
)

// Node represents a node in the Red-Black Tree
type Node struct {
	Key    int
	Value  interface{}
	Color  Color
	Left   *Node
	Right  *Node
	Parent *Node
}

// RBTree represents a Red-Black Tree
type RBTree struct {
	Root *Node
	Size int
}

// NewRBTree creates a new Red-Black Tree
func NewRBTree() *RBTree {
	return &RBTree{
		Root: nil,
		Size: 0,
	}
}

// Insert adds a new key-value pair to the tree
func (t *RBTree) Insert(key int, value interface{}) {
	newNode := &Node{
		Key:   key,
		Value: value,
		Color: Red,
	}

	if t.Root == nil {
		t.Root = newNode
		t.Root.Color = Black
		t.Size++
		return
	}

	// Standard BST insertion
	current := t.Root
	var parent *Node

	for current != nil {
		parent = current
		if key < current.Key {
			current = current.Left
		} else if key > current.Key {
			current = current.Right
		} else {
			// Key already exists, update value
			current.Value = value
			return
		}
	}

	newNode.Parent = parent
	if key < parent.Key {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	t.Size++
	t.fixInsert(newNode)
}

// fixInsert fixes the Red-Black Tree properties after insertion
func (t *RBTree) fixInsert(node *Node) {
	for node != t.Root && node.Parent.Color == Red {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == Red {
				// Case 1: Uncle is red
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					// Case 2: Node is right child
					node = node.Parent
					t.rotateLeft(node)
				}
				// Case 3: Node is left child
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				t.rotateRight(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == Red {
				// Case 1: Uncle is red
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					// Case 2: Node is left child
					node = node.Parent
					t.rotateRight(node)
				}
				// Case 3: Node is right child
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				t.rotateLeft(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = Black
}

// rotateLeft performs a left rotation on the given node
func (t *RBTree) rotateLeft(node *Node) {
	rightChild := node.Right
	node.Right = rightChild.Left

	if rightChild.Left != nil {
		rightChild.Left.Parent = node
	}

	rightChild.Parent = node.Parent

	if node.Parent == nil {
		t.Root = rightChild
	} else if node == node.Parent.Left {
		node.Parent.Left = rightChild
	} else {
		node.Parent.Right = rightChild
	}

	rightChild.Left = node
	node.Parent = rightChild
}

// rotateRight performs a right rotation on the given node
func (t *RBTree) rotateRight(node *Node) {
	leftChild := node.Left
	node.Left = leftChild.Right

	if leftChild.Right != nil {
		leftChild.Right.Parent = node
	}

	leftChild.Parent = node.Parent

	if node.Parent == nil {
		t.Root = leftChild
	} else if node == node.Parent.Right {
		node.Parent.Right = leftChild
	} else {
		node.Parent.Left = leftChild
	}

	leftChild.Right = node
	node.Parent = leftChild
}

// Search finds a node with the given key
func (t *RBTree) Search(key int) (interface{}, bool) {
	node := t.searchNode(key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

// searchNode returns the node with the given key
func (t *RBTree) searchNode(key int) *Node {
	current := t.Root
	for current != nil {
		if key == current.Key {
			return current
		} else if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return nil
}

// Delete removes a node with the given key from the tree
func (t *RBTree) Delete(key int) bool {
	node := t.searchNode(key)
	if node == nil {
		return false
	}

	t.deleteNode(node)
	t.Size--
	return true
}

// deleteNode removes the given node from the tree
func (t *RBTree) deleteNode(node *Node) {
	var replacement *Node
	originalColor := node.Color

	if node.Left == nil {
		replacement = node.Right
		t.transplant(node, node.Right)
	} else if node.Right == nil {
		replacement = node.Left
		t.transplant(node, node.Left)
	} else {
		// Node has two children
		successor := t.minimum(node.Right)
		originalColor = successor.Color
		replacement = successor.Right

		if successor.Parent == node {
			if replacement != nil {
				replacement.Parent = successor
			}
		} else {
			t.transplant(successor, successor.Right)
			successor.Right = node.Right
			successor.Right.Parent = successor
		}

		t.transplant(node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
		successor.Color = node.Color
	}

	if originalColor == Black && replacement != nil {
		t.fixDelete(replacement)
	}
}

// transplant replaces one subtree with another
func (t *RBTree) transplant(u, v *Node) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

// minimum finds the minimum node in a subtree
func (t *RBTree) minimum(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// fixDelete fixes the Red-Black Tree properties after deletion
func (t *RBTree) fixDelete(node *Node) {
	for node != t.Root && node.Color == Black {
		if node == node.Parent.Left {
			sibling := node.Parent.Right
			if sibling.Color == Red {
				// Case 1: Sibling is red
				sibling.Color = Black
				node.Parent.Color = Red
				t.rotateLeft(node.Parent)
				sibling = node.Parent.Right
			}

			if (sibling.Left == nil || sibling.Left.Color == Black) &&
				(sibling.Right == nil || sibling.Right.Color == Black) {
				// Case 2: Sibling's children are black
				sibling.Color = Red
				node = node.Parent
			} else {
				if sibling.Right == nil || sibling.Right.Color == Black {
					// Case 3: Sibling's right child is black
					if sibling.Left != nil {
						sibling.Left.Color = Black
					}
					sibling.Color = Red
					t.rotateRight(sibling)
					sibling = node.Parent.Right
				}

				// Case 4: Sibling's right child is red
				sibling.Color = node.Parent.Color
				node.Parent.Color = Black
				if sibling.Right != nil {
					sibling.Right.Color = Black
				}
				t.rotateLeft(node.Parent)
				node = t.Root
			}
		} else {
			sibling := node.Parent.Left
			if sibling.Color == Red {
				// Case 1: Sibling is red
				sibling.Color = Black
				node.Parent.Color = Red
				t.rotateRight(node.Parent)
				sibling = node.Parent.Left
			}

			if (sibling.Right == nil || sibling.Right.Color == Black) &&
				(sibling.Left == nil || sibling.Left.Color == Black) {
				// Case 2: Sibling's children are black
				sibling.Color = Red
				node = node.Parent
			} else {
				if sibling.Left == nil || sibling.Left.Color == Black {
					// Case 3: Sibling's left child is black
					if sibling.Right != nil {
						sibling.Right.Color = Black
					}
					sibling.Color = Red
					t.rotateLeft(sibling)
					sibling = node.Parent.Left
				}

				// Case 4: Sibling's left child is red
				sibling.Color = node.Parent.Color
				node.Parent.Color = Black
				if sibling.Left != nil {
					sibling.Left.Color = Black
				}
				t.rotateRight(node.Parent)
				node = t.Root
			}
		}
	}
	node.Color = Black
}

// InOrderTraversal performs an in-order traversal of the tree
func (t *RBTree) InOrderTraversal(visit func(key int, value interface{})) {
	t.inOrderHelper(t.Root, visit)
}

func (t *RBTree) inOrderHelper(node *Node, visit func(key int, value interface{})) {
	if node != nil {
		t.inOrderHelper(node.Left, visit)
		visit(node.Key, node.Value)
		t.inOrderHelper(node.Right, visit)
	}
}

// GetSize returns the number of nodes in the tree
func (t *RBTree) GetSize() int {
	return t.Size
}

// IsEmpty checks if the tree is empty
func (t *RBTree) IsEmpty() bool {
	return t.Size == 0
}

// Clear removes all nodes from the tree
func (t *RBTree) Clear() {
	t.Root = nil
	t.Size = 0
}

// String returns a string representation of the tree
func (t *RBTree) String() string {
	if t.Root == nil {
		return "Empty Tree"
	}
	return t.stringHelper(t.Root, "", true)
}

func (t *RBTree) stringHelper(node *Node, prefix string, isTail bool) string {
	if node == nil {
		return ""
	}

	result := ""
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		result += t.stringHelper(node.Right, newPrefix, false)
	}

	result += prefix
	if isTail {
		result += "└── "
	} else {
		result += "┌── "
	}

	color := "B"
	if node.Color == Red {
		color = "R"
	}
	result += fmt.Sprintf("%d(%s)\n", node.Key, color)

	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		result += t.stringHelper(node.Left, newPrefix, true)
	}

	return result
}

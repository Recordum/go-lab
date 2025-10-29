# B-Tree Implementation

This package provides a generic B-tree implementation in Go.

## Overview

A B-tree is a self-balancing tree data structure that maintains sorted data and allows searches, sequential access, insertions, and deletions in logarithmic time. B-trees are particularly well-suited for storage systems that read and write large blocks of data.

## Features

- **Generic implementation** using Go generics (supports any ordered type)
- **Insert operation** - Add keys to the tree with automatic balancing
- **Search operation** - Find keys in O(log n) time
- **Delete operation** - Remove keys with automatic rebalancing
- **Traverse operation** - In-order traversal of all keys
- **Configurable degree** - Set the minimum degree (t) for the tree

## Properties

- All leaves are at the same depth
- A node can contain multiple keys (up to 2t-1 where t is the minimum degree)
- All nodes except root must contain at least t-1 keys
- All nodes can contain at most 2t-1 keys
- Keys in each node are stored in sorted order

## Usage

### Creating a B-Tree

```go
package main

import (
    "fmt"
    "your-module/data-structure/btree"
)

func main() {
    // Create a B-tree with minimum degree 3
    bt := btree.NewBTree[int](3)
}
```

### Inserting Keys

```go
bt.Insert(10)
bt.Insert(20)
bt.Insert(5)
bt.Insert(6)
bt.Insert(12)
bt.Insert(30)
```

### Searching for Keys

```go
found := bt.Search(20)
if found {
    fmt.Println("Key 20 exists in the tree")
}
```

### Deleting Keys

```go
deleted := bt.Delete(10)
if deleted {
    fmt.Println("Key 10 was successfully deleted")
}
```

### Traversing the Tree

```go
keys := bt.Traverse()
fmt.Println("Keys in sorted order:", keys)
```

### Getting Tree Height

```go
height := bt.Height()
fmt.Printf("Tree height: %d\n", height)
```

### Printing Tree Structure (Debug)

```go
bt.Print()
```

## Example

```go
package main

import (
    "fmt"
    "your-module/data-structure/btree"
)

func main() {
    // Create a B-tree with minimum degree 3
    bt := btree.NewBTree[int](3)

    // Insert keys
    keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
    for _, key := range keys {
        bt.Insert(key)
        fmt.Printf("Inserted %d\n", key)
    }

    // Print tree structure
    fmt.Println("\nTree structure:")
    bt.Print()

    // Traverse and print keys
    fmt.Println("\nKeys in sorted order:")
    fmt.Println(bt.Traverse())

    // Search for keys
    fmt.Println("\nSearching for keys:")
    for _, key := range []int{6, 15, 17} {
        found := bt.Search(key)
        fmt.Printf("Key %d: %v\n", key, found)
    }

    // Delete keys
    fmt.Println("\nDeleting keys:")
    for _, key := range []int{6, 7, 20} {
        deleted := bt.Delete(key)
        fmt.Printf("Delete %d: %v\n", key, deleted)
    }

    // Print final state
    fmt.Println("\nFinal tree:")
    fmt.Println(bt.Traverse())
}
```

## Time Complexity

- **Search**: O(log n)
- **Insert**: O(log n)
- **Delete**: O(log n)
- **Traverse**: O(n)

where n is the number of keys in the tree.

## Space Complexity

O(n) where n is the number of keys in the tree.

## Testing

Run the tests with:

```bash
go test ./data-structure/btree
```

Run benchmarks with:

```bash
go test -bench=. ./data-structure/btree
```

## References

- [Introduction to Algorithms (CLRS)](https://mitpress.mit.edu/books/introduction-algorithms-third-edition) - Chapter 18: B-Trees
- [Wikipedia: B-tree](https://en.wikipedia.org/wiki/B-tree)

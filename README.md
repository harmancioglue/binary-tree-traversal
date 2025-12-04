# Binary Tree Traversal Implementations in Go

<img width="800" height="400" alt="image" src="https://github.com/user-attachments/assets/d5e6d734-694f-4cc8-8a07-cd591a52b99a" />

# Binary Tree Traversal Implementations in Go

This repository contains implementations of Binary Tree traversal algorithms in Go.

## Contents

Four fundamental methods for Binary Tree traversal are implemented:

### 1. Preorder Traversal
**Order:** Root → Left → Right

The root node is visited first, then the left subtree, and finally the right subtree.

**Use Cases:**
- Tree copying
- Creating prefix expressions

### 2. Inorder Traversal
**Order:** Left → Root → Right

The left subtree is visited first, then the root node, and finally the right subtree.

**Use Cases:**
- Getting sorted list from Binary Search Tree
- Printing BST elements in ascending order

### 3. Postorder Traversal
**Order:** Left → Right → Root

Both subtrees are traversed first, and the root node is visited last.

**Use Cases:**
- Tree deletion operations
- Calculating file system size

### 4. Level Order Traversal (BFS)
**Order:** Level by level, left to right

The tree is traversed level by level, from left to right at each level.

**Use Cases:**
- Finding shortest path
- Calculating minimum depth
- Breadth-First Search (BFS)

## Example
```
        1
       / \
      2   3
     / \
    4   5
```

**Outputs:**
- Preorder: `[1, 2, 4, 5, 3]`
- Inorder: `[4, 2, 5, 1, 3]`
- Postorder: `[4, 5, 2, 3, 1]`
- Level Order: `[1, 2, 3, 4, 5]`

## Implementation Details

Both **recursive** and **iterative** implementations are provided for each traversal method.

- **Preorder, Inorder, Postorder:** Using Stack (DFS - Depth First Search)
- **Level Order:** Using Queue (BFS - Breadth First Search)

## Complexity

- **Time Complexity:** O(n) - Each node is visited once
- **Space Complexity:** O(h) - h: tree height (for recursive stack)

## Resources

- [LeetCode - Binary Tree Traversal Problems](https://leetcode.com/)
- [GeeksforGeeks - Tree Traversals](https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/)


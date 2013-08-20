// To Do
// -------
// Add deletion
// Add traversal

package main

import "fmt"

type Node struct {
  right, left *Node
  value int
}

type BinaryTree struct {
  root *Node
}

func addNode(val int) *Node {
  return &Node{nil, nil, val}
}

func (b *BinaryTree) Insert(val int) {
  if b.root == nil {
    b.root = addNode(val)
  } else {
    b.insert(b.root, val)
  }
}

func (b *BinaryTree) insert(root *Node, val int) *Node {
  if root == nil {
    return addNode(val)
  }
  if val <= root.value {
    root.left = b.insert(root.left, val)
  } else {
    root.right = b.insert(root.right, val)
  }
  return root
}

func main() {
  b := new(BinaryTree)
  fmt.Println(b.root)
  b.Insert(0)
  b.Insert(-1)
  b.Insert(1)
  fmt.Println(b.root)
  fmt.Println(b.root.right)
  fmt.Println(b.root.left)
  b.Insert(2)
  fmt.Println(b.root.right.right)
}
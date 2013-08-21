// To Do
// -------
// Add deletion
// Add traversal

package main

import "fmt"

type Node struct {
  right, left, parent *Node
  value int
}

type BinaryTree struct {
  root *Node
}

func addNode(parent *Node, val int) *Node {
  return &Node{nil, nil, parent, val}
}

func (b *BinaryTree) Insert(val int) (n *Node) {
  if b.root == nil {
    n = addNode(nil, val)
    b.root = n
  } else {
    n = b.insert(b.root, nil, val)
  }
  return
}

func (b *BinaryTree) insert(root *Node, parent *Node, val int) *Node {
  if root == nil {
    return addNode(parent, val)
  }
  if val <= root.value {
    root.left = b.insert(root.left, root, val)
  } else {
    root.right = b.insert(root.right, root, val)
  }
  return root
}

func main() {
  b := new(BinaryTree)
  fmt.Println(b.root)
  b.Insert(0)
  fmt.Println(b.root)
  b.Insert(1)
  b.Insert(-1)
  fmt.Println(b.root)
  fmt.Println(b.root.right)
  fmt.Println(b.root.left)
}
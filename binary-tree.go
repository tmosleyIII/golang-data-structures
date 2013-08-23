// To Do
// -------
// Make delete a little cleaner
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
  switch {
  case root == nil:
    return addNode(parent, val)
  case val <= root.value:
    root.left = b.insert(root.left, root, val)
  case val > root.value:
    root.right = b.insert(root.right, root, val)
  }
  return root
}

func (b *BinaryTree) Delete(val int) {
  if b.root == nil {
    return
  } else {
    b.delete(b.root, val)
  }
}

func (b *BinaryTree) delete (root *Node, val int) (n *Node) {
  if b.root.value == val && b.root.left == nil && b.root.right == nil {
    n = b.root
    b.root = nil
    return
  } else if root.value != val {
    if val <= root.value {
      b.delete(root.left, val)
    } else if val > root.value {
      b.delete(root.right, val)
    }
  } else if root.value == val && root.left == nil && root.right == nil {
    n = b.root
    b.root = nil
    return
  } else if root.value == val && (root.left == nil || root.right == nil) {
    // there's a better, cleaner way to do this in one line
    var n1 *Node
    if root.left != nil {
      n1 = root.left
    } else {
      n1 = root.right
    }
    switch {
    case root.parent.left == root:
      root.parent.left = n1
      return root
    case root.parent.right == root:
      root.parent.right = n1
      return root
    } 
  } else if root.value == val && root.left != nil && root.right != nil {
    // Needs to be implemented later
  }
  return nil
}

func main() {
  b := new(BinaryTree)
  b.Insert(0)
  b.Insert(1)
  b.Insert(-1)
  fmt.Println(b.root)
  b.Delete(1)
  fmt.Println(b.root)
}
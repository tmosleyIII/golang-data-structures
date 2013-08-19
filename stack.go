package main

import (
  "fmt"
)

type Stack struct {
  top *Node
  size int
}

type Node struct {
  value string
  next *Node
}

func (s *Stack) Length() int {
  return s.size
}

func (s *Stack) IsEmpty() bool {
  if s.size == 0 {
    return true
  }
  return false
}

func (s *Stack) Push(val string) {
  s.top = &Node{val, s.top}
  s.size++
}

func (s *Stack) Peek() (val string) {
  return s.top.value
}

func (s *Stack) Pop() (val string) {
  if s.size > 0 {
    val, s.top = s.top.value, s.top.next
    s.size--
    return val
  }
  return ""
}

func main() {
  stack := new(Stack)

  stack.Push("Orr")
  stack.Push("It")
  stack.Push("Ship")
  stack.Push("Scott")
  stack.Push("Michael")
  stack.Push("Zachary")

  for stack.Length() > 0 {
    fmt.Printf("%s ", stack.Pop())
  }
  fmt.Println()
}
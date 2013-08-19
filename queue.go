package main

import (
  "fmt"
)

type Queue struct {
  first *Node
  size int
}

type Node struct {
  value string
  next *Node
}

func (q *Queue) Length() int {
  return q.size
}

func (q *Queue) IsEmpty() bool {
  return q.size == 0
}

func (q *Queue) Enqueue(val string) {
  if q.first == nil {
    q.first = &Node{val, nil}
  } else {
    last := q.first
    for last.next != nil {
      last = last.next
    }
    last.next = &Node{val, nil}
  }
  q.size++
}

func (q *Queue) Peek() string {
  return q.first.value
}

func (q *Queue) Dequeue() (val string) {
  if q.size > 0 {
    val, q.first = q.first.value, q.first.next
    q.size--
    return
  }
  return ""
}

func main() {
  queue := new(Queue)

  queue.Enqueue("Zachary")
  queue.Enqueue("Michael")
  queue.Enqueue("Scott")
  queue.Enqueue("Ship")
  queue.Enqueue("It")
  queue.Enqueue("Orr")

  for queue.Length() > 0 {
    fmt.Printf("%s ", queue.Dequeue())
  }
  fmt.Println()
}
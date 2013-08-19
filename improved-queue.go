package main

import (
  "fmt"
)

type Queue struct {
  first, last *Node
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
  n := &Node{val, nil}
  if q.first == nil {
    q.first, q.last = n, n
  } else {
    last := q.last
    last.next = n
    q.last = n
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
    if q.size == 0 {
      q.last = nil
    }
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
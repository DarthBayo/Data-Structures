package main

import "fmt"

type QNode struct {
  value any
  next *QNode
}

type Queue struct {
  head, tail *QNode
  length uint8
}

type QueueInterface interface {
  dequeue() *QNode
  enqueue(v any)
}

func (q *Queue) enqueue(v any) {
  node := &QNode{
    value: v,
    next: nil,
  }

  q.length++
  if q.tail == nil {
    q.head, q.tail = node, node
    return
  }

  q.tail.next = node
  q.tail = node
}

func (q *Queue) dequeue() {
  if q.head == nil {
    return
  }
  q.length--

  node := q.head
  q.head = node.next
  node = nil

  if q.head == nil {
    q.tail = nil
  }
}

func main() {
  q := &Queue{}

  q.enqueue(1)
  q.enqueue(2)
  q.enqueue(3)
  q.enqueue(4)

  q.dequeue()
  q.dequeue()
  q.dequeue()
  q.dequeue()
  q.dequeue()

  fmt.Printf("%+v\n", q)
  fmt.Printf("%d\n", q.head)
  fmt.Printf("%d\n", q.tail)
}

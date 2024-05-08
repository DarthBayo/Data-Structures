package main

import "fmt"

type SNode struct {
	value any
	next  *SNode
}

type Stack struct {
	top    *SNode
	length int
}

type StackInterface interface {
	push(v any)
	pop()
}

func (s *Stack) push(v any) {
	node := &SNode{
		value: v,
		next:  nil,
	}

	s.length++
	if s.top == nil {
		s.top = node
		return
	}

	node.next = s.top
	s.top = node
}

func (s *Stack) pop() {
	if s.top == nil {
		return
	}
	s.length--

	s.top = s.top.next
}

func main() {
	s := &Stack{}

	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)

	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()

	fmt.Printf("%+v\n", s)
	fmt.Printf("%d\n", s.top)
}
